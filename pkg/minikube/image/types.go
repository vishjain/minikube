/*
Copyright 2020 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image

// ErrNeedsLogin is thrown when registry needs login (a general error)
var ErrNeedsLogin error

// ErrGithubNeedsLogin is thrown when user needs to login specifically to github packages)
var ErrGithubNeedsLogin error

// ErrFailedDownloadKICImages is thrown when user
var ErrFailedDownloadKICImages error
