package legal

// ===========================================================================================================
// This file was generated automatically at 15-11-2021 23:46:01 using gogenlicense.
// Do not edit manually, as changes may be overwritten.
// ===========================================================================================================

// The following go packages are imported:
// - Go Standard Library (BSD-3-Clause; see https://golang.org/LICENSE)
// - gopkg.in/warnings.v0 (BSD-2-Clause)
// - gopkg.in/src-d/go-git.v4 (Apache-2.0)
// - gopkg.in/src-d/go-billy.v4 (Apache-2.0)
// - golang.org/x/xerrors (BSD-3-Clause)
// - golang.org/x/tools (BSD-3-Clause)
// - golang.org/x/sys (BSD-3-Clause)
// - golang.org/x/net (BSD-3-Clause)
// - golang.org/x/mod (BSD-3-Clause)
// - golang.org/x/crypto (BSD-3-Clause)
// - github.com/xanzy/ssh-agent (Apache-2.0; see https://github.com/xanzy/ssh-agent/blob/master/LICENSE)
// - github.com/src-d/gcfg (BSD-3-Clause; see https://github.com/src-d/gcfg/blob/master/LICENSE)
// - github.com/sergi/go-diff/diffmatchpatch (MIT; see https://github.com/sergi/go-diff/blob/master/diffmatchpatch/LICENSE)
// - github.com/pkg/errors (BSD-2-Clause; see https://github.com/pkg/errors/blob/master/LICENSE)
// - github.com/mitchellh/go-homedir (MIT; see https://github.com/mitchellh/go-homedir/blob/master/LICENSE)
// - github.com/kevinburke/ssh_config (MIT; see https://github.com/kevinburke/ssh_config/blob/master/LICENSE)
// - github.com/jbenet/go-context/io (MIT; see https://github.com/jbenet/go-context/blob/master/io/LICENSE)
// - github.com/google/licenseclassifier/stringclassifier (Apache-2.0; see https://github.com/google/licenseclassifier/blob/master/stringclassifier/LICENSE)
// - github.com/google/licenseclassifier (Apache-2.0; see https://github.com/google/licenseclassifier/blob/master/LICENSE)
// - github.com/google/go-licenses/licenses (Apache-2.0; see https://github.com/google/go-licenses/blob/master/licenses/LICENSE)
// - github.com/golang/glog (Apache-2.0; see https://github.com/golang/glog/blob/master/LICENSE)
// - github.com/emirpasic/gods (BSD-2-Clause; see https://github.com/emirpasic/gods/blob/master/LICENSE)
//
// ================================================================================
//
//
// ================================================================================
// Go Standard Library
// Licensed under the Terms of the BSD-3-Clause License, see also https://golang.org/LICENSE.
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// 	* Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// 	* Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
// 	* Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module gopkg.in/warnings.v0
// Licensed under the Terms of the BSD-2-Clause License
//
// Copyright (c) 2016 Péter Surányi.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module gopkg.in/src-d/go-git.v4
// Licensed under the Terms of the Apache-2.0 License
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "{}"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright 2018 Sourced Technologies, S.L.
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
//
// ================================================================================
//
// ================================================================================
// Module gopkg.in/src-d/go-billy.v4
// Licensed under the Terms of the Apache-2.0 License
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "{}"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright 2017 Sourced Technologies S.L.
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
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/xerrors
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2019 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/tools
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/sys
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/net
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/mod
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module golang.org/x/crypto
// Licensed under the Terms of the BSD-3-Clause License
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/xanzy/ssh-agent
// Licensed under the Terms of the Apache-2.0 License, see also https://github.com/xanzy/ssh-agent/blob/master/LICENSE.
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "{}"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright {yyyy} {name of copyright owner}
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
//
//
// ================================================================================
//
// ================================================================================
// Module github.com/src-d/gcfg
// Licensed under the Terms of the BSD-3-Clause License, see also https://github.com/src-d/gcfg/blob/master/LICENSE.
//
// Copyright (c) 2012 Péter Surányi. Portions Copyright (c) 2009 The Go
// Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/sergi/go-diff/diffmatchpatch
// Licensed under the Terms of the MIT License, see also https://github.com/sergi/go-diff/blob/master/diffmatchpatch/LICENSE.
//
// Copyright (c) 2012-2016 The go-diff Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.
//
//
// ================================================================================
//
// ================================================================================
// Module github.com/pkg/errors
// Licensed under the Terms of the BSD-2-Clause License, see also https://github.com/pkg/errors/blob/master/LICENSE.
//
// Copyright (c) 2015, Dave Cheney <dave@cheney.net>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/mitchellh/go-homedir
// Licensed under the Terms of the MIT License, see also https://github.com/mitchellh/go-homedir/blob/master/LICENSE.
//
// The MIT License (MIT)
//
// Copyright (c) 2013 Mitchell Hashimoto
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/kevinburke/ssh_config
// Licensed under the Terms of the MIT License, see also https://github.com/kevinburke/ssh_config/blob/master/LICENSE.
//
// Copyright (c) 2017 Kevin Burke.
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// ===================
//
// The lexer and parser borrow heavily from github.com/pelletier/go-toml. The
// license for that project is copied below.
//
// The MIT License (MIT)
//
// Copyright (c) 2013 - 2017 Thomas Pelletier, Eric Anderton
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/jbenet/go-context/io
// Licensed under the Terms of the MIT License, see also https://github.com/jbenet/go-context/blob/master/io/LICENSE.
//
// The MIT License (MIT)
//
// Copyright (c) 2014 Juan Batiz-Benet
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//
// ================================================================================
//
// ================================================================================
// Module github.com/google/licenseclassifier/stringclassifier
// Licensed under the Terms of the Apache-2.0 License, see also https://github.com/google/licenseclassifier/blob/master/stringclassifier/LICENSE.
//
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "[]"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright [yyyy] [name of copyright owner]
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
//
// ================================================================================
//
// ================================================================================
// Module github.com/google/licenseclassifier
// Licensed under the Terms of the Apache-2.0 License, see also https://github.com/google/licenseclassifier/blob/master/LICENSE.
//
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "[]"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright [yyyy] [name of copyright owner]
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
//
// ================================================================================
//
// ================================================================================
// Module github.com/google/go-licenses/licenses
// Licensed under the Terms of the Apache-2.0 License, see also https://github.com/google/go-licenses/blob/master/licenses/LICENSE.
//
//
//                                  Apache License
//                            Version 2.0, January 2004
//                         http://www.apache.org/licenses/
//
//    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
//    1. Definitions.
//
//       "License" shall mean the terms and conditions for use, reproduction,
//       and distribution as defined by Sections 1 through 9 of this document.
//
//       "Licensor" shall mean the copyright owner or entity authorized by
//       the copyright owner that is granting the License.
//
//       "Legal Entity" shall mean the union of the acting entity and all
//       other entities that control, are controlled by, or are under common
//       control with that entity. For the purposes of this definition,
//       "control" means (i) the power, direct or indirect, to cause the
//       direction or management of such entity, whether by contract or
//       otherwise, or (ii) ownership of fifty percent (50%) or more of the
//       outstanding shares, or (iii) beneficial ownership of such entity.
//
//       "You" (or "Your") shall mean an individual or Legal Entity
//       exercising permissions granted by this License.
//
//       "Source" form shall mean the preferred form for making modifications,
//       including but not limited to software source code, documentation
//       source, and configuration files.
//
//       "Object" form shall mean any form resulting from mechanical
//       transformation or translation of a Source form, including but
//       not limited to compiled object code, generated documentation,
//       and conversions to other media types.
//
//       "Work" shall mean the work of authorship, whether in Source or
//       Object form, made available under the License, as indicated by a
//       copyright notice that is included in or attached to the work
//       (an example is provided in the Appendix below).
//
//       "Derivative Works" shall mean any work, whether in Source or Object
//       form, that is based on (or derived from) the Work and for which the
//       editorial revisions, annotations, elaborations, or other modifications
//       represent, as a whole, an original work of authorship. For the purposes
//       of this License, Derivative Works shall not include works that remain
//       separable from, or merely link (or bind by name) to the interfaces of,
//       the Work and Derivative Works thereof.
//
//       "Contribution" shall mean any work of authorship, including
//       the original version of the Work and any modifications or additions
//       to that Work or Derivative Works thereof, that is intentionally
//       submitted to Licensor for inclusion in the Work by the copyright owner
//       or by an individual or Legal Entity authorized to submit on behalf of
//       the copyright owner. For the purposes of this definition, "submitted"
//       means any form of electronic, verbal, or written communication sent
//       to the Licensor or its representatives, including but not limited to
//       communication on electronic mailing lists, source code control systems,
//       and issue tracking systems that are managed by, or on behalf of, the
//       Licensor for the purpose of discussing and improving the Work, but
//       excluding communication that is conspicuously marked or otherwise
//       designated in writing by the copyright owner as "Not a Contribution."
//
//       "Contributor" shall mean Licensor and any individual or Legal Entity
//       on behalf of whom a Contribution has been received by Licensor and
//       subsequently incorporated within the Work.
//
//    2. Grant of Copyright License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       copyright license to reproduce, prepare Derivative Works of,
//       publicly display, publicly perform, sublicense, and distribute the
//       Work and such Derivative Works in Source or Object form.
//
//    3. Grant of Patent License. Subject to the terms and conditions of
//       this License, each Contributor hereby grants to You a perpetual,
//       worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//       (except as stated in this section) patent license to make, have made,
//       use, offer to sell, sell, import, and otherwise transfer the Work,
//       where such license applies only to those patent claims licensable
//       by such Contributor that are necessarily infringed by their
//       Contribution(s) alone or by combination of their Contribution(s)
//       with the Work to which such Contribution(s) was submitted. If You
//       institute patent litigation against any entity (including a
//       cross-claim or counterclaim in a lawsuit) alleging that the Work
//       or a Contribution incorporated within the Work constitutes direct
//       or contributory patent infringement, then any patent licenses
//       granted to You under this License for that Work shall terminate
//       as of the date such litigation is filed.
//
//    4. Redistribution. You may reproduce and distribute copies of the
//       Work or Derivative Works thereof in any medium, with or without
//       modifications, and in Source or Object form, provided that You
//       meet the following conditions:
//
//       (a) You must give any other recipients of the Work or
//           Derivative Works a copy of this License; and
//
//       (b) You must cause any modified files to carry prominent notices
//           stating that You changed the files; and
//
//       (c) You must retain, in the Source form of any Derivative Works
//           that You distribute, all copyright, patent, trademark, and
//           attribution notices from the Source form of the Work,
//           excluding those notices that do not pertain to any part of
//           the Derivative Works; and
//
//       (d) If the Work includes a "NOTICE" text file as part of its
//           distribution, then any Derivative Works that You distribute must
//           include a readable copy of the attribution notices contained
//           within such NOTICE file, excluding those notices that do not
//           pertain to any part of the Derivative Works, in at least one
//           of the following places: within a NOTICE text file distributed
//           as part of the Derivative Works; within the Source form or
//           documentation, if provided along with the Derivative Works; or,
//           within a display generated by the Derivative Works, if and
//           wherever such third-party notices normally appear. The contents
//           of the NOTICE file are for informational purposes only and
//           do not modify the License. You may add Your own attribution
//           notices within Derivative Works that You distribute, alongside
//           or as an addendum to the NOTICE text from the Work, provided
//           that such additional attribution notices cannot be construed
//           as modifying the License.
//
//       You may add Your own copyright statement to Your modifications and
//       may provide additional or different license terms and conditions
//       for use, reproduction, or distribution of Your modifications, or
//       for any such Derivative Works as a whole, provided Your use,
//       reproduction, and distribution of the Work otherwise complies with
//       the conditions stated in this License.
//
//    5. Submission of Contributions. Unless You explicitly state otherwise,
//       any Contribution intentionally submitted for inclusion in the Work
//       by You to the Licensor shall be under the terms and conditions of
//       this License, without any additional terms or conditions.
//       Notwithstanding the above, nothing herein shall supersede or modify
//       the terms of any separate license agreement you may have executed
//       with Licensor regarding such Contributions.
//
//    6. Trademarks. This License does not grant permission to use the trade
//       names, trademarks, service marks, or product names of the Licensor,
//       except as required for reasonable and customary use in describing the
//       origin of the Work and reproducing the content of the NOTICE file.
//
//    7. Disclaimer of Warranty. Unless required by applicable law or
//       agreed to in writing, Licensor provides the Work (and each
//       Contributor provides its Contributions) on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//       implied, including, without limitation, any warranties or conditions
//       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//       PARTICULAR PURPOSE. You are solely responsible for determining the
//       appropriateness of using or redistributing the Work and assume any
//       risks associated with Your exercise of permissions under this License.
//
//    8. Limitation of Liability. In no event and under no legal theory,
//       whether in tort (including negligence), contract, or otherwise,
//       unless required by applicable law (such as deliberate and grossly
//       negligent acts) or agreed to in writing, shall any Contributor be
//       liable to You for damages, including any direct, indirect, special,
//       incidental, or consequential damages of any character arising as a
//       result of this License or out of the use or inability to use the
//       Work (including but not limited to damages for loss of goodwill,
//       work stoppage, computer failure or malfunction, or any and all
//       other commercial damages or losses), even if such Contributor
//       has been advised of the possibility of such damages.
//
//    9. Accepting Warranty or Additional Liability. While redistributing
//       the Work or Derivative Works thereof, You may choose to offer,
//       and charge a fee for, acceptance of support, warranty, indemnity,
//       or other liability obligations and/or rights consistent with this
//       License. However, in accepting such obligations, You may act only
//       on Your own behalf and on Your sole responsibility, not on behalf
//       of any other Contributor, and only if You agree to indemnify,
//       defend, and hold each Contributor harmless for any liability
//       incurred by, or claims asserted against, such Contributor by reason
//       of your accepting any such warranty or additional liability.
//
//    END OF TERMS AND CONDITIONS
//
//    APPENDIX: How to apply the Apache License to your work.
//
//       To apply the Apache License to your work, attach the following
//       boilerplate notice, with the fields enclosed by brackets "[]"
//       replaced with your own identifying information. (Don't include
//       the brackets!)  The text should be enclosed in the appropriate
//       comment syntax for the file format. We also recommend that a
//       file or class name and description of purpose be included on the
//       same "printed page" as the copyright notice for easier
//       identification within third-party archives.
//
//    Copyright [yyyy] [name of copyright owner]
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
//
// ================================================================================
//
// ================================================================================
// Module github.com/golang/glog
// Licensed under the Terms of the Apache-2.0 License, see also https://github.com/golang/glog/blob/master/LICENSE.
//
// Apache License
// Version 2.0, January 2004
// http://www.apache.org/licenses/
//
// TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
// 1. Definitions.
//
// "License" shall mean the terms and conditions for use, reproduction, and
// distribution as defined by Sections 1 through 9 of this document.
//
// "Licensor" shall mean the copyright owner or entity authorized by the copyright
// owner that is granting the License.
//
// "Legal Entity" shall mean the union of the acting entity and all other entities
// that control, are controlled by, or are under common control with that entity.
// For the purposes of this definition, "control" means (i) the power, direct or
// indirect, to cause the direction or management of such entity, whether by
// contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the
// outstanding shares, or (iii) beneficial ownership of such entity.
//
// "You" (or "Your") shall mean an individual or Legal Entity exercising
// permissions granted by this License.
//
// "Source" form shall mean the preferred form for making modifications, including
// but not limited to software source code, documentation source, and configuration
// files.
//
// "Object" form shall mean any form resulting from mechanical transformation or
// translation of a Source form, including but not limited to compiled object code,
// generated documentation, and conversions to other media types.
//
// "Work" shall mean the work of authorship, whether in Source or Object form, made
// available under the License, as indicated by a copyright notice that is included
// in or attached to the work (an example is provided in the Appendix below).
//
// "Derivative Works" shall mean any work, whether in Source or Object form, that
// is based on (or derived from) the Work and for which the editorial revisions,
// annotations, elaborations, or other modifications represent, as a whole, an
// original work of authorship. For the purposes of this License, Derivative Works
// shall not include works that remain separable from, or merely link (or bind by
// name) to the interfaces of, the Work and Derivative Works thereof.
//
// "Contribution" shall mean any work of authorship, including the original version
// of the Work and any modifications or additions to that Work or Derivative Works
// thereof, that is intentionally submitted to Licensor for inclusion in the Work
// by the copyright owner or by an individual or Legal Entity authorized to submit
// on behalf of the copyright owner. For the purposes of this definition,
// "submitted" means any form of electronic, verbal, or written communication sent
// to the Licensor or its representatives, including but not limited to
// communication on electronic mailing lists, source code control systems, and
// issue tracking systems that are managed by, or on behalf of, the Licensor for
// the purpose of discussing and improving the Work, but excluding communication
// that is conspicuously marked or otherwise designated in writing by the copyright
// owner as "Not a Contribution."
//
// "Contributor" shall mean Licensor and any individual or Legal Entity on behalf
// of whom a Contribution has been received by Licensor and subsequently
// incorporated within the Work.
//
// 2. Grant of Copyright License.
//
// Subject to the terms and conditions of this License, each Contributor hereby
// grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,
// irrevocable copyright license to reproduce, prepare Derivative Works of,
// publicly display, publicly perform, sublicense, and distribute the Work and such
// Derivative Works in Source or Object form.
//
// 3. Grant of Patent License.
//
// Subject to the terms and conditions of this License, each Contributor hereby
// grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,
// irrevocable (except as stated in this section) patent license to make, have
// made, use, offer to sell, sell, import, and otherwise transfer the Work, where
// such license applies only to those patent claims licensable by such Contributor
// that are necessarily infringed by their Contribution(s) alone or by combination
// of their Contribution(s) with the Work to which such Contribution(s) was
// submitted. If You institute patent litigation against any entity (including a
// cross-claim or counterclaim in a lawsuit) alleging that the Work or a
// Contribution incorporated within the Work constitutes direct or contributory
// patent infringement, then any patent licenses granted to You under this License
// for that Work shall terminate as of the date such litigation is filed.
//
// 4. Redistribution.
//
// You may reproduce and distribute copies of the Work or Derivative Works thereof
// in any medium, with or without modifications, and in Source or Object form,
// provided that You meet the following conditions:
//
// You must give any other recipients of the Work or Derivative Works a copy of
// this License; and
// You must cause any modified files to carry prominent notices stating that You
// changed the files; and
// You must retain, in the Source form of any Derivative Works that You distribute,
// all copyright, patent, trademark, and attribution notices from the Source form
// of the Work, excluding those notices that do not pertain to any part of the
// Derivative Works; and
// If the Work includes a "NOTICE" text file as part of its distribution, then any
// Derivative Works that You distribute must include a readable copy of the
// attribution notices contained within such NOTICE file, excluding those notices
// that do not pertain to any part of the Derivative Works, in at least one of the
// following places: within a NOTICE text file distributed as part of the
// Derivative Works; within the Source form or documentation, if provided along
// with the Derivative Works; or, within a display generated by the Derivative
// Works, if and wherever such third-party notices normally appear. The contents of
// the NOTICE file are for informational purposes only and do not modify the
// License. You may add Your own attribution notices within Derivative Works that
// You distribute, alongside or as an addendum to the NOTICE text from the Work,
// provided that such additional attribution notices cannot be construed as
// modifying the License.
// You may add Your own copyright statement to Your modifications and may provide
// additional or different license terms and conditions for use, reproduction, or
// distribution of Your modifications, or for any such Derivative Works as a whole,
// provided Your use, reproduction, and distribution of the Work otherwise complies
// with the conditions stated in this License.
//
// 5. Submission of Contributions.
//
// Unless You explicitly state otherwise, any Contribution intentionally submitted
// for inclusion in the Work by You to the Licensor shall be under the terms and
// conditions of this License, without any additional terms or conditions.
// Notwithstanding the above, nothing herein shall supersede or modify the terms of
// any separate license agreement you may have executed with Licensor regarding
// such Contributions.
//
// 6. Trademarks.
//
// This License does not grant permission to use the trade names, trademarks,
// service marks, or product names of the Licensor, except as required for
// reasonable and customary use in describing the origin of the Work and
// reproducing the content of the NOTICE file.
//
// 7. Disclaimer of Warranty.
//
// Unless required by applicable law or agreed to in writing, Licensor provides the
// Work (and each Contributor provides its Contributions) on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied,
// including, without limitation, any warranties or conditions of TITLE,
// NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are
// solely responsible for determining the appropriateness of using or
// redistributing the Work and assume any risks associated with Your exercise of
// permissions under this License.
//
// 8. Limitation of Liability.
//
// In no event and under no legal theory, whether in tort (including negligence),
// contract, or otherwise, unless required by applicable law (such as deliberate
// and grossly negligent acts) or agreed to in writing, shall any Contributor be
// liable to You for damages, including any direct, indirect, special, incidental,
// or consequential damages of any character arising as a result of this License or
// out of the use or inability to use the Work (including but not limited to
// damages for loss of goodwill, work stoppage, computer failure or malfunction, or
// any and all other commercial damages or losses), even if such Contributor has
// been advised of the possibility of such damages.
//
// 9. Accepting Warranty or Additional Liability.
//
// While redistributing the Work or Derivative Works thereof, You may choose to
// offer, and charge a fee for, acceptance of support, warranty, indemnity, or
// other liability obligations and/or rights consistent with this License. However,
// in accepting such obligations, You may act only on Your own behalf and on Your
// sole responsibility, not on behalf of any other Contributor, and only if You
// agree to indemnify, defend, and hold each Contributor harmless for any liability
// incurred by, or claims asserted against, such Contributor by reason of your
// accepting any such warranty or additional liability.
//
// END OF TERMS AND CONDITIONS
//
// APPENDIX: How to apply the Apache License to your work
//
// To apply the Apache License to your work, attach the following boilerplate
// notice, with the fields enclosed by brackets "[]" replaced with your own
// identifying information. (Don't include the brackets!) The text should be
// enclosed in the appropriate comment syntax for the file format. We also
// recommend that a file or class name and description of purpose be included on
// the same "printed page" as the copyright notice for easier identification within
// third-party archives.
//
//    Copyright [yyyy] [name of copyright owner]
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//
// ================================================================================
//
// ================================================================================
// Module github.com/emirpasic/gods
// Licensed under the Terms of the BSD-2-Clause License, see also https://github.com/emirpasic/gods/blob/master/LICENSE.
//
// Copyright (c) 2015, Emir Pasic
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// -------------------------------------------------------------------------------
//
// AVL Tree:
//
// Copyright (c) 2017 Benjamin Scher Purcell <benjapurcell@gmail.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//
// ================================================================================
//
var Notices string

func init() {
	Notices = "Notices contains legal and license information of external software included in this program. \nThese notices consist of a list of dependencies along with their license information.\nThis string is intended to be displayed to the enduser on demand.\n\nEven though the value of this variable is fixed at compile time it is omitted from this documentation.\nInstead the list of go modules, along with their licenses, is listed below. \n\n\nGo Standard Library\n\nThe Go Standard Library is licensed under the Terms of the BSD-3-Clause License. \nSee also https://golang.org/LICENSE. \n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n \t* Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n \t* Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n \t* Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule gopkg in warnings v0\n\nThe Module gopkg.in/warnings.v0 is licensed under the Terms of the BSD-2-Clause License. \n\n\n Copyright (c) 2016 Péter Surányi.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule gopkg in src d go git v4\n\nThe Module gopkg.in/src-d/go-git.v4 is licensed under the Terms of the Apache-2.0 License. \n\n\n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"{}\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright 2018 Sourced Technologies, S.L.\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule gopkg in src d go billy v4\n\nThe Module gopkg.in/src-d/go-billy.v4 is licensed under the Terms of the Apache-2.0 License. \n\n\n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"{}\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright 2017 Sourced Technologies S.L.\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule golang org x xerrors\n\nThe Module golang.org/x/xerrors is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2019 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule golang org x tools\n\nThe Module golang.org/x/tools is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule golang org x sys\n\nThe Module golang.org/x/sys is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule golang org x net\n\nThe Module golang.org/x/net is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule golang org x mod\n\nThe Module golang.org/x/mod is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule golang org x crypto\n\nThe Module golang.org/x/crypto is licensed under the Terms of the BSD-3-Clause License. \n\n\n Copyright (c) 2009 The Go Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule github com xanzy ssh agent\n\nThe Module github.com/xanzy/ssh-agent is licensed under the Terms of the Apache-2.0 License. \nSee also https://github.com/xanzy/ssh-agent/blob/master/LICENSE. \n\n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"{}\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright {yyyy} {name of copyright owner}\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n \n\nModule github com src d gcfg\n\nThe Module github.com/src-d/gcfg is licensed under the Terms of the BSD-3-Clause License. \nSee also https://github.com/src-d/gcfg/blob/master/LICENSE. \n\n Copyright (c) 2012 Péter Surányi. Portions Copyright (c) 2009 The Go\n Authors. All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are\n met:\n \n    * Redistributions of source code must retain the above copyright\n notice, this list of conditions and the following disclaimer.\n    * Redistributions in binary form must reproduce the above\n copyright notice, this list of conditions and the following disclaimer\n in the documentation and/or other materials provided with the\n distribution.\n    * Neither the name of Google Inc. nor the names of its\n contributors may be used to endorse or promote products derived from\n this software without specific prior written permission.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS\n \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT\n LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR\n A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT\n OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,\n SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT\n LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,\n DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY\n THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule github com sergi go diff diffmatchpatch\n\nThe Module github.com/sergi/go-diff/diffmatchpatch is licensed under the Terms of the MIT License. \nSee also https://github.com/sergi/go-diff/blob/master/diffmatchpatch/LICENSE. \n\n Copyright (c) 2012-2016 The go-diff Authors. All rights reserved.\n \n Permission is hereby granted, free of charge, to any person obtaining a\n copy of this software and associated documentation files (the \"Software\"),\n to deal in the Software without restriction, including without limitation\n the rights to use, copy, modify, merge, publish, distribute, sublicense,\n and/or sell copies of the Software, and to permit persons to whom the\n Software is furnished to do so, subject to the following conditions:\n \n The above copyright notice and this permission notice shall be included\n in all copies or substantial portions of the Software.\n \n THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS\n OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING\n FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER\n DEALINGS IN THE SOFTWARE.\n \n \n\nModule github com pkg errors\n\nThe Module github.com/pkg/errors is licensed under the Terms of the BSD-2-Clause License. \nSee also https://github.com/pkg/errors/blob/master/LICENSE. \n\n Copyright (c) 2015, Dave Cheney <dave@cheney.net>\n All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are met:\n \n * Redistributions of source code must retain the above copyright notice, this\n   list of conditions and the following disclaimer.\n \n * Redistributions in binary form must reproduce the above copyright notice,\n   this list of conditions and the following disclaimer in the documentation\n   and/or other materials provided with the distribution.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\"\n AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE\n IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\n DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\n FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\n DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\n SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\n CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\n OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n\nModule github com mitchellh go homedir\n\nThe Module github.com/mitchellh/go-homedir is licensed under the Terms of the MIT License. \nSee also https://github.com/mitchellh/go-homedir/blob/master/LICENSE. \n\n The MIT License (MIT)\n \n Copyright (c) 2013 Mitchell Hashimoto\n \n Permission is hereby granted, free of charge, to any person obtaining a copy\n of this software and associated documentation files (the \"Software\"), to deal\n in the Software without restriction, including without limitation the rights\n to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n copies of the Software, and to permit persons to whom the Software is\n furnished to do so, subject to the following conditions:\n \n The above copyright notice and this permission notice shall be included in\n all copies or substantial portions of the Software.\n \n THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n THE SOFTWARE.\n \n\nModule github com kevinburke ssh config\n\nThe Module github.com/kevinburke/ssh_config is licensed under the Terms of the MIT License. \nSee also https://github.com/kevinburke/ssh_config/blob/master/LICENSE. \n\n Copyright (c) 2017 Kevin Burke.\n \n Permission is hereby granted, free of charge, to any person\n obtaining a copy of this software and associated documentation\n files (the \"Software\"), to deal in the Software without\n restriction, including without limitation the rights to use,\n copy, modify, merge, publish, distribute, sublicense, and/or sell\n copies of the Software, and to permit persons to whom the\n Software is furnished to do so, subject to the following\n conditions:\n \n The above copyright notice and this permission notice shall be\n included in all copies or substantial portions of the Software.\n \n THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND,\n EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES\n OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND\n NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT\n HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,\n WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING\n FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR\n OTHER DEALINGS IN THE SOFTWARE.\n \n ===================\n \n The lexer and parser borrow heavily from github.com/pelletier/go-toml. The\n license for that project is copied below.\n \n The MIT License (MIT)\n \n Copyright (c) 2013 - 2017 Thomas Pelletier, Eric Anderton\n \n Permission is hereby granted, free of charge, to any person obtaining a copy\n of this software and associated documentation files (the \"Software\"), to deal\n in the Software without restriction, including without limitation the rights\n to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n copies of the Software, and to permit persons to whom the Software is\n furnished to do so, subject to the following conditions:\n \n The above copyright notice and this permission notice shall be included in all\n copies or substantial portions of the Software.\n \n THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\n SOFTWARE.\n \n\nModule github com jbenet go context io\n\nThe Module github.com/jbenet/go-context/io is licensed under the Terms of the MIT License. \nSee also https://github.com/jbenet/go-context/blob/master/io/LICENSE. \n\n The MIT License (MIT)\n \n Copyright (c) 2014 Juan Batiz-Benet\n \n Permission is hereby granted, free of charge, to any person obtaining a copy\n of this software and associated documentation files (the \"Software\"), to deal\n in the Software without restriction, including without limitation the rights\n to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n copies of the Software, and to permit persons to whom the Software is\n furnished to do so, subject to the following conditions:\n \n The above copyright notice and this permission notice shall be included in\n all copies or substantial portions of the Software.\n \n THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n THE SOFTWARE.\n \n\nModule github com google licenseclassifier stringclassifier\n\nThe Module github.com/google/licenseclassifier/stringclassifier is licensed under the Terms of the Apache-2.0 License. \nSee also https://github.com/google/licenseclassifier/blob/master/stringclassifier/LICENSE. \n\n \n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"[]\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright [yyyy] [name of copyright owner]\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule github com google licenseclassifier\n\nThe Module github.com/google/licenseclassifier is licensed under the Terms of the Apache-2.0 License. \nSee also https://github.com/google/licenseclassifier/blob/master/LICENSE. \n\n \n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"[]\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright [yyyy] [name of copyright owner]\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule github com google go licenses licenses\n\nThe Module github.com/google/go-licenses/licenses is licensed under the Terms of the Apache-2.0 License. \nSee also https://github.com/google/go-licenses/blob/master/licenses/LICENSE. \n\n \n                                  Apache License\n                            Version 2.0, January 2004\n                         http://www.apache.org/licenses/\n \n    TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n    1. Definitions.\n \n       \"License\" shall mean the terms and conditions for use, reproduction,\n       and distribution as defined by Sections 1 through 9 of this document.\n \n       \"Licensor\" shall mean the copyright owner or entity authorized by\n       the copyright owner that is granting the License.\n \n       \"Legal Entity\" shall mean the union of the acting entity and all\n       other entities that control, are controlled by, or are under common\n       control with that entity. For the purposes of this definition,\n       \"control\" means (i) the power, direct or indirect, to cause the\n       direction or management of such entity, whether by contract or\n       otherwise, or (ii) ownership of fifty percent (50%) or more of the\n       outstanding shares, or (iii) beneficial ownership of such entity.\n \n       \"You\" (or \"Your\") shall mean an individual or Legal Entity\n       exercising permissions granted by this License.\n \n       \"Source\" form shall mean the preferred form for making modifications,\n       including but not limited to software source code, documentation\n       source, and configuration files.\n \n       \"Object\" form shall mean any form resulting from mechanical\n       transformation or translation of a Source form, including but\n       not limited to compiled object code, generated documentation,\n       and conversions to other media types.\n \n       \"Work\" shall mean the work of authorship, whether in Source or\n       Object form, made available under the License, as indicated by a\n       copyright notice that is included in or attached to the work\n       (an example is provided in the Appendix below).\n \n       \"Derivative Works\" shall mean any work, whether in Source or Object\n       form, that is based on (or derived from) the Work and for which the\n       editorial revisions, annotations, elaborations, or other modifications\n       represent, as a whole, an original work of authorship. For the purposes\n       of this License, Derivative Works shall not include works that remain\n       separable from, or merely link (or bind by name) to the interfaces of,\n       the Work and Derivative Works thereof.\n \n       \"Contribution\" shall mean any work of authorship, including\n       the original version of the Work and any modifications or additions\n       to that Work or Derivative Works thereof, that is intentionally\n       submitted to Licensor for inclusion in the Work by the copyright owner\n       or by an individual or Legal Entity authorized to submit on behalf of\n       the copyright owner. For the purposes of this definition, \"submitted\"\n       means any form of electronic, verbal, or written communication sent\n       to the Licensor or its representatives, including but not limited to\n       communication on electronic mailing lists, source code control systems,\n       and issue tracking systems that are managed by, or on behalf of, the\n       Licensor for the purpose of discussing and improving the Work, but\n       excluding communication that is conspicuously marked or otherwise\n       designated in writing by the copyright owner as \"Not a Contribution.\"\n \n       \"Contributor\" shall mean Licensor and any individual or Legal Entity\n       on behalf of whom a Contribution has been received by Licensor and\n       subsequently incorporated within the Work.\n \n    2. Grant of Copyright License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       copyright license to reproduce, prepare Derivative Works of,\n       publicly display, publicly perform, sublicense, and distribute the\n       Work and such Derivative Works in Source or Object form.\n \n    3. Grant of Patent License. Subject to the terms and conditions of\n       this License, each Contributor hereby grants to You a perpetual,\n       worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n       (except as stated in this section) patent license to make, have made,\n       use, offer to sell, sell, import, and otherwise transfer the Work,\n       where such license applies only to those patent claims licensable\n       by such Contributor that are necessarily infringed by their\n       Contribution(s) alone or by combination of their Contribution(s)\n       with the Work to which such Contribution(s) was submitted. If You\n       institute patent litigation against any entity (including a\n       cross-claim or counterclaim in a lawsuit) alleging that the Work\n       or a Contribution incorporated within the Work constitutes direct\n       or contributory patent infringement, then any patent licenses\n       granted to You under this License for that Work shall terminate\n       as of the date such litigation is filed.\n \n    4. Redistribution. You may reproduce and distribute copies of the\n       Work or Derivative Works thereof in any medium, with or without\n       modifications, and in Source or Object form, provided that You\n       meet the following conditions:\n \n       (a) You must give any other recipients of the Work or\n           Derivative Works a copy of this License; and\n \n       (b) You must cause any modified files to carry prominent notices\n           stating that You changed the files; and\n \n       (c) You must retain, in the Source form of any Derivative Works\n           that You distribute, all copyright, patent, trademark, and\n           attribution notices from the Source form of the Work,\n           excluding those notices that do not pertain to any part of\n           the Derivative Works; and\n \n       (d) If the Work includes a \"NOTICE\" text file as part of its\n           distribution, then any Derivative Works that You distribute must\n           include a readable copy of the attribution notices contained\n           within such NOTICE file, excluding those notices that do not\n           pertain to any part of the Derivative Works, in at least one\n           of the following places: within a NOTICE text file distributed\n           as part of the Derivative Works; within the Source form or\n           documentation, if provided along with the Derivative Works; or,\n           within a display generated by the Derivative Works, if and\n           wherever such third-party notices normally appear. The contents\n           of the NOTICE file are for informational purposes only and\n           do not modify the License. You may add Your own attribution\n           notices within Derivative Works that You distribute, alongside\n           or as an addendum to the NOTICE text from the Work, provided\n           that such additional attribution notices cannot be construed\n           as modifying the License.\n \n       You may add Your own copyright statement to Your modifications and\n       may provide additional or different license terms and conditions\n       for use, reproduction, or distribution of Your modifications, or\n       for any such Derivative Works as a whole, provided Your use,\n       reproduction, and distribution of the Work otherwise complies with\n       the conditions stated in this License.\n \n    5. Submission of Contributions. Unless You explicitly state otherwise,\n       any Contribution intentionally submitted for inclusion in the Work\n       by You to the Licensor shall be under the terms and conditions of\n       this License, without any additional terms or conditions.\n       Notwithstanding the above, nothing herein shall supersede or modify\n       the terms of any separate license agreement you may have executed\n       with Licensor regarding such Contributions.\n \n    6. Trademarks. This License does not grant permission to use the trade\n       names, trademarks, service marks, or product names of the Licensor,\n       except as required for reasonable and customary use in describing the\n       origin of the Work and reproducing the content of the NOTICE file.\n \n    7. Disclaimer of Warranty. Unless required by applicable law or\n       agreed to in writing, Licensor provides the Work (and each\n       Contributor provides its Contributions) on an \"AS IS\" BASIS,\n       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n       implied, including, without limitation, any warranties or conditions\n       of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n       PARTICULAR PURPOSE. You are solely responsible for determining the\n       appropriateness of using or redistributing the Work and assume any\n       risks associated with Your exercise of permissions under this License.\n \n    8. Limitation of Liability. In no event and under no legal theory,\n       whether in tort (including negligence), contract, or otherwise,\n       unless required by applicable law (such as deliberate and grossly\n       negligent acts) or agreed to in writing, shall any Contributor be\n       liable to You for damages, including any direct, indirect, special,\n       incidental, or consequential damages of any character arising as a\n       result of this License or out of the use or inability to use the\n       Work (including but not limited to damages for loss of goodwill,\n       work stoppage, computer failure or malfunction, or any and all\n       other commercial damages or losses), even if such Contributor\n       has been advised of the possibility of such damages.\n \n    9. Accepting Warranty or Additional Liability. While redistributing\n       the Work or Derivative Works thereof, You may choose to offer,\n       and charge a fee for, acceptance of support, warranty, indemnity,\n       or other liability obligations and/or rights consistent with this\n       License. However, in accepting such obligations, You may act only\n       on Your own behalf and on Your sole responsibility, not on behalf\n       of any other Contributor, and only if You agree to indemnify,\n       defend, and hold each Contributor harmless for any liability\n       incurred by, or claims asserted against, such Contributor by reason\n       of your accepting any such warranty or additional liability.\n \n    END OF TERMS AND CONDITIONS\n \n    APPENDIX: How to apply the Apache License to your work.\n \n       To apply the Apache License to your work, attach the following\n       boilerplate notice, with the fields enclosed by brackets \"[]\"\n       replaced with your own identifying information. (Don't include\n       the brackets!)  The text should be enclosed in the appropriate\n       comment syntax for the file format. We also recommend that a\n       file or class name and description of purpose be included on the\n       same \"printed page\" as the copyright notice for easier\n       identification within third-party archives.\n \n    Copyright [yyyy] [name of copyright owner]\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n        http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule github com golang glog\n\nThe Module github.com/golang/glog is licensed under the Terms of the Apache-2.0 License. \nSee also https://github.com/golang/glog/blob/master/LICENSE. \n\n Apache License\n Version 2.0, January 2004\n http://www.apache.org/licenses/\n \n TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n \n 1. Definitions.\n \n \"License\" shall mean the terms and conditions for use, reproduction, and\n distribution as defined by Sections 1 through 9 of this document.\n \n \"Licensor\" shall mean the copyright owner or entity authorized by the copyright\n owner that is granting the License.\n \n \"Legal Entity\" shall mean the union of the acting entity and all other entities\n that control, are controlled by, or are under common control with that entity.\n For the purposes of this definition, \"control\" means (i) the power, direct or\n indirect, to cause the direction or management of such entity, whether by\n contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the\n outstanding shares, or (iii) beneficial ownership of such entity.\n \n \"You\" (or \"Your\") shall mean an individual or Legal Entity exercising\n permissions granted by this License.\n \n \"Source\" form shall mean the preferred form for making modifications, including\n but not limited to software source code, documentation source, and configuration\n files.\n \n \"Object\" form shall mean any form resulting from mechanical transformation or\n translation of a Source form, including but not limited to compiled object code,\n generated documentation, and conversions to other media types.\n \n \"Work\" shall mean the work of authorship, whether in Source or Object form, made\n available under the License, as indicated by a copyright notice that is included\n in or attached to the work (an example is provided in the Appendix below).\n \n \"Derivative Works\" shall mean any work, whether in Source or Object form, that\n is based on (or derived from) the Work and for which the editorial revisions,\n annotations, elaborations, or other modifications represent, as a whole, an\n original work of authorship. For the purposes of this License, Derivative Works\n shall not include works that remain separable from, or merely link (or bind by\n name) to the interfaces of, the Work and Derivative Works thereof.\n \n \"Contribution\" shall mean any work of authorship, including the original version\n of the Work and any modifications or additions to that Work or Derivative Works\n thereof, that is intentionally submitted to Licensor for inclusion in the Work\n by the copyright owner or by an individual or Legal Entity authorized to submit\n on behalf of the copyright owner. For the purposes of this definition,\n \"submitted\" means any form of electronic, verbal, or written communication sent\n to the Licensor or its representatives, including but not limited to\n communication on electronic mailing lists, source code control systems, and\n issue tracking systems that are managed by, or on behalf of, the Licensor for\n the purpose of discussing and improving the Work, but excluding communication\n that is conspicuously marked or otherwise designated in writing by the copyright\n owner as \"Not a Contribution.\"\n \n \"Contributor\" shall mean Licensor and any individual or Legal Entity on behalf\n of whom a Contribution has been received by Licensor and subsequently\n incorporated within the Work.\n \n 2. Grant of Copyright License.\n \n Subject to the terms and conditions of this License, each Contributor hereby\n grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,\n irrevocable copyright license to reproduce, prepare Derivative Works of,\n publicly display, publicly perform, sublicense, and distribute the Work and such\n Derivative Works in Source or Object form.\n \n 3. Grant of Patent License.\n \n Subject to the terms and conditions of this License, each Contributor hereby\n grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,\n irrevocable (except as stated in this section) patent license to make, have\n made, use, offer to sell, sell, import, and otherwise transfer the Work, where\n such license applies only to those patent claims licensable by such Contributor\n that are necessarily infringed by their Contribution(s) alone or by combination\n of their Contribution(s) with the Work to which such Contribution(s) was\n submitted. If You institute patent litigation against any entity (including a\n cross-claim or counterclaim in a lawsuit) alleging that the Work or a\n Contribution incorporated within the Work constitutes direct or contributory\n patent infringement, then any patent licenses granted to You under this License\n for that Work shall terminate as of the date such litigation is filed.\n \n 4. Redistribution.\n \n You may reproduce and distribute copies of the Work or Derivative Works thereof\n in any medium, with or without modifications, and in Source or Object form,\n provided that You meet the following conditions:\n \n You must give any other recipients of the Work or Derivative Works a copy of\n this License; and\n You must cause any modified files to carry prominent notices stating that You\n changed the files; and\n You must retain, in the Source form of any Derivative Works that You distribute,\n all copyright, patent, trademark, and attribution notices from the Source form\n of the Work, excluding those notices that do not pertain to any part of the\n Derivative Works; and\n If the Work includes a \"NOTICE\" text file as part of its distribution, then any\n Derivative Works that You distribute must include a readable copy of the\n attribution notices contained within such NOTICE file, excluding those notices\n that do not pertain to any part of the Derivative Works, in at least one of the\n following places: within a NOTICE text file distributed as part of the\n Derivative Works; within the Source form or documentation, if provided along\n with the Derivative Works; or, within a display generated by the Derivative\n Works, if and wherever such third-party notices normally appear. The contents of\n the NOTICE file are for informational purposes only and do not modify the\n License. You may add Your own attribution notices within Derivative Works that\n You distribute, alongside or as an addendum to the NOTICE text from the Work,\n provided that such additional attribution notices cannot be construed as\n modifying the License.\n You may add Your own copyright statement to Your modifications and may provide\n additional or different license terms and conditions for use, reproduction, or\n distribution of Your modifications, or for any such Derivative Works as a whole,\n provided Your use, reproduction, and distribution of the Work otherwise complies\n with the conditions stated in this License.\n \n 5. Submission of Contributions.\n \n Unless You explicitly state otherwise, any Contribution intentionally submitted\n for inclusion in the Work by You to the Licensor shall be under the terms and\n conditions of this License, without any additional terms or conditions.\n Notwithstanding the above, nothing herein shall supersede or modify the terms of\n any separate license agreement you may have executed with Licensor regarding\n such Contributions.\n \n 6. Trademarks.\n \n This License does not grant permission to use the trade names, trademarks,\n service marks, or product names of the Licensor, except as required for\n reasonable and customary use in describing the origin of the Work and\n reproducing the content of the NOTICE file.\n \n 7. Disclaimer of Warranty.\n \n Unless required by applicable law or agreed to in writing, Licensor provides the\n Work (and each Contributor provides its Contributions) on an \"AS IS\" BASIS,\n WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied,\n including, without limitation, any warranties or conditions of TITLE,\n NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are\n solely responsible for determining the appropriateness of using or\n redistributing the Work and assume any risks associated with Your exercise of\n permissions under this License.\n \n 8. Limitation of Liability.\n \n In no event and under no legal theory, whether in tort (including negligence),\n contract, or otherwise, unless required by applicable law (such as deliberate\n and grossly negligent acts) or agreed to in writing, shall any Contributor be\n liable to You for damages, including any direct, indirect, special, incidental,\n or consequential damages of any character arising as a result of this License or\n out of the use or inability to use the Work (including but not limited to\n damages for loss of goodwill, work stoppage, computer failure or malfunction, or\n any and all other commercial damages or losses), even if such Contributor has\n been advised of the possibility of such damages.\n \n 9. Accepting Warranty or Additional Liability.\n \n While redistributing the Work or Derivative Works thereof, You may choose to\n offer, and charge a fee for, acceptance of support, warranty, indemnity, or\n other liability obligations and/or rights consistent with this License. However,\n in accepting such obligations, You may act only on Your own behalf and on Your\n sole responsibility, not on behalf of any other Contributor, and only if You\n agree to indemnify, defend, and hold each Contributor harmless for any liability\n incurred by, or claims asserted against, such Contributor by reason of your\n accepting any such warranty or additional liability.\n \n END OF TERMS AND CONDITIONS\n \n APPENDIX: How to apply the Apache License to your work\n \n To apply the Apache License to your work, attach the following boilerplate\n notice, with the fields enclosed by brackets \"[]\" replaced with your own\n identifying information. (Don't include the brackets!) The text should be\n enclosed in the appropriate comment syntax for the file format. We also\n recommend that a file or class name and description of purpose be included on\n the same \"printed page\" as the copyright notice for easier identification within\n third-party archives.\n \n    Copyright [yyyy] [name of copyright owner]\n \n    Licensed under the Apache License, Version 2.0 (the \"License\");\n    you may not use this file except in compliance with the License.\n    You may obtain a copy of the License at\n \n      http://www.apache.org/licenses/LICENSE-2.0\n \n    Unless required by applicable law or agreed to in writing, software\n    distributed under the License is distributed on an \"AS IS\" BASIS,\n    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n    See the License for the specific language governing permissions and\n    limitations under the License.\n \n\nModule github com emirpasic gods\n\nThe Module github.com/emirpasic/gods is licensed under the Terms of the BSD-2-Clause License. \nSee also https://github.com/emirpasic/gods/blob/master/LICENSE. \n\n Copyright (c) 2015, Emir Pasic\n All rights reserved.\n \n Redistribution and use in source and binary forms, with or without\n modification, are permitted provided that the following conditions are met:\n \n * Redistributions of source code must retain the above copyright notice, this\n   list of conditions and the following disclaimer.\n \n * Redistributions in binary form must reproduce the above copyright notice,\n   this list of conditions and the following disclaimer in the documentation\n   and/or other materials provided with the distribution.\n \n THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\"\n AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE\n IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\n DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\n FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\n DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\n SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\n CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\n OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\n OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n \n -------------------------------------------------------------------------------\n \n AVL Tree:\n \n Copyright (c) 2017 Benjamin Scher Purcell <benjapurcell@gmail.com>\n \n Permission to use, copy, modify, and distribute this software for any\n purpose with or without fee is hereby granted, provided that the above\n copyright notice and this permission notice appear in all copies.\n \n THE SOFTWARE IS PROVIDED \"AS IS\" AND THE AUTHOR DISCLAIMS ALL WARRANTIES\n WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF\n MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR\n ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES\n WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN\n ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF\n OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.\n \n\n\nGeneration\n\nThis variable and the associated documentation have been automatically generated using the 'gogenlicense' tool. \nIt was last updated at 15-11-2021 23:46:01. \n"
}
