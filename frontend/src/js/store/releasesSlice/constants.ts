// Copyright 2019 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
export const ARTIFACT_GENERATION_TYPE = { SINGLE_FILE: 'single_file' };

export const currentArtifact = 'artifact_name';
export const softwareIndicator = '.version';
export const rootfsImageVersion = 'rootfs-image.version';
export const softwareTitleMap = { [rootfsImageVersion]: { title: 'Root filesystem', priority: 0, key: rootfsImageVersion } };
