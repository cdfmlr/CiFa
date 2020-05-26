# CiFa

![CiFa-logo](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf5ntapw3nj30wi0fsgoh.jpg)

>  CiFa 是一个**词频统计**工具，提供一套 Web API。后端使用 Golang 实现，完成计算密集型任务，并维持 http 服务。前端基于 Vue.js，实现了一套 Ant Design 风格的用户界面。

该仓库储存 CiFa 的后端代码，前端仓库转：[https://github.com/cdfmlr/CiFa-front](https://github.com/cdfmlr/CiFa-front)。


## 功能概览

- 以《数据结构与算法》课程中所学的“串”结构为基础，实现了如下字符串模式匹配的算法：
   1. 暴力匹配算法
   2. KMP算法
   3. Rabin-Karp 算法
- 以《数据结构与算法》课程中所学的“排序”算法为基础，实现了排序算法：
   1. 快速排序
   2. 堆排序
   3. 归并排序
   4. 希尔排序
   5. 插入排序
   6. 选择排序
- 在“字符串模式匹配算法”与“排序算法”实现的基础上，实现了词频统计工具。即给定关键词，选择另外一个文本文件，程序会统计关键词在该文件中出现的频数，完毕后以频数从大到小的顺序进行输出。或者也可以选择一个目录打包成的 zip 文件，统计关键词在该目录下所有文本文件中出现的频数。

## Getting Started

- For **Linux / MacOS**:

```sh
$ wget https://github.com/cdfmlr/CiFa/blob/master/install.sh
$ bash install.sh
```

- For **Windows**:

WSL (recommended) or building from source manually.

- From **Source**:

Require: `git`, `npm`, `go>=12`

```sh
git clone https://github.com/cdfmlr/CiFa.git	# clone back end
git clone https://github.com/cdfmlr/CiFa-front.git	# clone front end

# build back end
cd CiFa/main
go build main.go
# You will get the bin file: `main`
main --help		# see usages
cd ../..

# build front end
cd CiFa-front
npm install -g @vue/cli		# install vue cli if not exist
npm install ant-design-vue
npm run build	# web ui (static ) will be built to ./dist
cd ..
```

## 开发进度

- [x] 排序算法
- [x] 字符串搜索算法
- [x] 词频统计实现
- [x] 服务框架
- [x] 词频统计 API、web UI
- [x] 搜索 API、web UI
- [x] 排序 API、web UI
- [ ] CLI

## 开放源代码

Copyright 2020 CDFMLR                                                   

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.






