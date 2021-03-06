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

import { StaticPage } from "./components/staticPage";
import { serve, IComponent } from "./utils/provider";

// add additional components to this array.
// make sure that components use the @PulumiComponent and @ComponentState decorators from "./utils/decorators"
const components: IComponent[] = [ StaticPage ];

// starts up the provider
serve(components);
