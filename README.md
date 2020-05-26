# CiFa

![CiFa-logo](https://tva1.sinaimg.cn/large/007S8ZIlgy1gf5ntapw3nj30wi0fsgoh.jpg)

>  CiFa 是一个**词频统计**工具，提供 Web API、Web GUI、CLI 三种用户接口。后端使用 Golang 完成 CLI 接口(基于 Cobra)、计算密集型任务(自己实现的算法)、http 服务(基于 net/http)。Web GUI (前端)基于 Vue.js，实现了一套 Ant Design 风格的用户界面。

该仓库储存 CiFa 的后端代码，前端仓库转：[https://github.com/cdfmlr/CiFa-front](https://github.com/cdfmlr/CiFa-front)。

## 功能概览

CiFa 主要是一个词频统计工具。

通过 [CiFa-front](https://github.com/cdfmlr/CiFa-front) 提供的 Web GUI （当然也能直接调用底层的 Web API），你可以给定一些关键词，并选择一个文本文件，程序将统计关键词在该文件中出现的频数，完毕后以频数从大到小的顺序进行输出。或者，你也可以选择一个目录打包成的 zip 文件，统计关键词在该目录下所有文本文件中出现的频数。

你还可以利用命令行工具，完成类似的操作：给定一个关键词文件，该文件存储所要统计词频的关键词。然后提供另外一个文本文件，统计关键词在该文件中出现的频数。或者可以选择一个目录，统计关键词在该目录下所有文本文件中出现的频数，统计完毕后以频数从大到小的顺序进行输出。

该词频统计工具基于以下算法实现：

- 以《数据结构与算法》课程中所学的“串”结构为基础，实现的字符串模式匹配的算法：
   1. 暴力匹配算法
   2. KMP算法
   3. Rabin-Karp 算法
- 以《数据结构与算法》课程中所学的“排序”算法为基础，实现的排序算法：
   1. 快速排序
   2. 堆排序
   3. 归并排序
   4. 希尔排序
   5. 插入排序
   6. 选择排序

## Getting Started

- For **Linux / MacOS**:

```sh
wget -N --no-check-certificate https://raw.githubusercontent.com/cdfmlr/CiFa/master/install.sh
bash install.sh
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

## 使用文档

### Web API

#### `wordfa`：词频统计接口

##### POST

>  `POST /api/wordfa`： 新建 wordfa 任务会话

- Request: 

```
POST /api/wordfa
```

- Request Form:

| key       | type             | description                                                  |
| --------- | ---------------- | ------------------------------------------------------------ |
| token     | FormValue string | 识别客户端身份的 token                                       |
| keywords  | FormValue string | 要检测的关键词，<br />多个词间用逗号(',' 或 '，')隔开        |
| file      | FormFile  file   | 要检测的文件，单个文本文件(text/plain)，<br />或多个文件的 zip 打包(application/zip) |
| sort_by   | FormValue int    | 结果的排序算法，0~8, 分别是：<br />sort.Sort (go lib)，sort.Stable (go lib)，快速排序，堆排序，归并排序，希尔排序，希尔排序(并发), 插入排序，选择排序 |
| search_by | FormValue int    | 字符串搜索算法，0~3                                          |

`sort_by` 是结果的排序算法，0~8 分别是：

| id | name      | description                       |
| --- | --------- | --------------------------------- |
| 0 | StlSort   | Go 标准库中的默认排序（快速排序） |
| 1 | StlStable | Go 标准库中的稳定排序（归并排序） |
| 2 | Quick     | 快速排序                          |
| 3 | Heap      | 堆排序                            |
| 4 | Merge     | 归并排序                          |
| 5 | Shell     | 希尔排序                          |
| 6 | ShellSync | 并发的希尔排序                    |
| 7 | Insertion | 插入排序                          |
| 8 | Selection | 选择排序                          |

`search_by` 字符串搜索算法，0~3 分别是：


| id | name      | description                        |
| --- | --------- | ---------------------------------- |
| 0 | LibRe     | 利用 Go 标准库中的正则表达式去匹配 |
| 1 | Kmp       | KMP 算法                           |
| 2 | RabinKarp | RabinKarp 算法                     |
| 3 | Naive     | 暴力算法                           |

- Response：

```
Success: JSON: {"success", "token"}
Failed:  JSON: {"error": "error description"}
```

##### GET

>  `GET /api/wordfa`： 获取 wordfa 会话任务的处理进度/结果

- Request:

```
GET /api/wordfa
```

- Request Form:

| key   | value            | description            |
| ----- | ---------------- | ---------------------- |
| token | FormValue string | 识别客户端身份的 token |

- Response:

```
Task Running:  JSON: {"progress": 0.7}
Task Finished: JSON: {"progress": 1.0, "result": [{"keyword": 26}, {...}, ...]}
Error:         JSON: {"error": "error description"}
```

#### `sort`：排序接口

> POST /api/sort/float, 对给定浮点数序列进行排序

- Request:

```
POST /api/sort/float
```

- RequestBody： JSON：

```json
{"algorithm": 0, "data": [2, 1, 3.0, 7, 4.4, ...]}
```

| key       | type    | description                                             |
| --------- | ------- | ------------------------------------------------------- |
| data      | []float | 要排序的数据                                            |
| algorithm | int     | 指定排序算法，0~8，同 wordfa POST 中对 `sort_by` 的说明 |

- Response：

```
Success: JSON: {"result": [1, 2, 3.0, 4.4, 7, ...], "time_cost": "time cost"}
Error:   JSON: {"error": "error description"}
```

#### `strsearch`：字符串匹配接口

> POST /api/strsearch, 在给定字符串做子串搜索

- Request:

```
POST /api/strsearch
```

- Request Body：JSON：

```json
{"algorithm": 0, "text": "abcbab", "pattern": "ab"}
```

| key       | type   | description                                                 |
| --------- | ------ | ----------------------------------------------------------- |
| text      | string | 父字符串，在此字符串中搜索子串 pattern                      |
| pattern   | string | 子字符串，在 text 中搜索此字符串                            |
| algorithm | int    | 字符串搜索算法，0~3，同 wordfa POST 中对 `search_by` 的说明 |

- Response：

```
Success: JSON: {"index": [0, 4], "time_cost": "time cost"}	// index 是 pattern 在 text 中出现位置的索引，注意中文字符不是"第几个字"！
Error:   JSON: {"error": "error description"}
```

### CLI

基本用法:

```
$ cifa [command]
```

可用的子命令:

| command | description                                 |
| ------- | ------------------------------------------- |
| serve   | Start a CiFa web serve                      |
| wordfa  | Run a words frequency analyzing task in CLI |
| help    | Help about any command                      |

#### cifa serve

`$ cifa serve` 可以开启一个 CiFa Web 服务，包括 Web API 服务和 Web GUI (CiFa-front) 的静态文件服务：

```sh
$ cifa serve
```

运行此命令后即可在 `http://localhost:9001` 访问 CiFa 服务。

（请使用 Getting Started 中的 install.sh 安装 CiFa，以得到正确的Web GUI CiFa-front 的静态文件服务地址，或者参考`cifa serve --help` 手动配置）

更多用法请看程序随附的命令行帮助：

```sh
$ cifa serve --help
```

#### cifa wordfa

`$ cifa wordfa` 在 CLI 中运行一个词频统计任务。

简单使用如下：

```
$ cifa wordfa -f test.txt -k keywords.txt -m Kmp -s Quick -o output.txt
```

这个命令就使用 `-m` 指定的 Kmp 字符串匹配算法、 `-s` 指定的 Quick 排序算法，在 `-f` 指定的文本文件 test.txt 检索 `-k`  指定的关键词文件 keywords.txt 中的所有关键词（可以以逗号或换行分隔），结果输出到 `-o` 指定的文件中，输出文件内容大致如下：

```
不是: 44254
可以: 26225
我的: 10378
你好: 804
再见: 392
```

可选的字符串匹配算法和排序算法参考 wordfa POST 部分的文档（在这里传入算法名称而不是id）。

更多用法请看程序随附的命令行帮助：

```sh
$ cifa wordfa --help
```



## 开发进度

- [x] 排序算法
- [x] 字符串搜索算法
- [x] 词频统计实现
- [x] 服务框架
- [x] 词频统计 API、web UI
- [x] 搜索 API、web UI
- [x] 排序 API、web UI
- [x] CLI

## 开放源代码

Copyright 2020 CDFMLR                                                   

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.






