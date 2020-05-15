// strsearch 包提供一系列字符串搜索算法
//
// algorithms implemented:
//  - NaiveSearchByChar
//  - NaiveSearchBySlice
//  - KmpSearch
//  - RabinKarpSearch
//
// Notes:
//  NaiveSearchBySlice is slower than NaiveSearchByChar
//  RabinKarpSearch is despised, for its md5 calling as a hash function, it's too slowwwwww.

package strsearch

/******************************************************************************
 *    Copyright 2020 CDFMLR                                                   *
 *                                                                            *
 *    Licensed under the Apache License, Version 2.0 (the "License");         *
 *    you may not use this file except in compliance with the License.        *
 *    You may obtain a copy of the License at                                 *
 *                                                                            *
 *        http://www.apache.org/licenses/LICENSE-2.0                          *
 *                                                                            *
 *    Unless required by applicable law or agreed to in writing, software     *
 *    distributed under the License is distributed on an "AS IS" BASIS,       *
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.*
 *    See the License for the specific language governing permissions and     *
 *    limitations under the License.                                          *
 *                                                                            *
 ******************************************************************************/
