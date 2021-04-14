// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";

export type IComponent = new (name: string, args: any, opts?: pulumi.ComponentResourceOptions) => pulumi.ComponentResource;

export class Provider implements provider.Provider {
    private tokenToComponent: Map<string, IComponent> = new Map();
    constructor(readonly version: string, components: IComponent[]) {
        for(const c of components){
            this.tokenToComponent.set(c.prototype.__typeToken, c)
        }
    }

    async construct(name: string, type: string, inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {
        let component = this.tokenToComponent.get(type);
        if (!component) {
            throw new Error(`unknown resource type ${type}`);
        }

        const instance = new component(name, inputs, options);
        const result: provider.ConstructResult = {
            urn: instance.urn,
            state: {},
        };

        for (let [k, _] of (instance as any).__state || new Map()) {
            result.state[k] = (instance as any)[k]
        }

        return result;
    }
}

export function serve(components: IComponent[]){
    function main(args: string[]) {
        let version: string = require("../package.json").version;
        // Node allows for the version to be prefixed by a "v",
        // while semver doesn't. If there is a v, strip it off.
        if (version.startsWith("v")) {
            version = version.slice(1);
        }
        return pulumi.provider.main(new Provider(version, components), args);
    }
    
    main(process.argv.slice(2));
}
