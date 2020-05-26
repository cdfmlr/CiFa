######################################################################
#
# FILE:   install.sh
# BY:     CDFMLR
# UPDATE: 2020.05.26
#
# Install Script for CiFa:
#     - back-end:  https://github.com/cdfmlr/CiFa.git
#     - front-end: https://github.com/cdfmlr/CiFa-front.git
#
# Copyright 2020 CDFMLR
# 
# Licensed under the Apache License, Version 2.0 (the "License"); 
# you may not use this file except in compliance with the License. 
# You may obtain a copy of the License at
# 
#        http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, 
# software distributed under the License is distributed on an 
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, 
# either express or implied. See the License for the specific 
# language governing permissions and limitations under the License.
#
#####################################################################

# 依赖检查
echo "👉 CiFa-install> 1/3 > Checke env..."

echo "git check: "
if command -v git; then
    echo "    Version `git --version` installed."
else
    echo "    git not found, please install it and try again."
    exit 1
fi

echo ""
echo "npm check: "
if command -v npm; then
    echo "    Version `npm -v` installed."
else
    echo "    npm not found, please install it and try again."
    exit 1
fi

echo ""
echo "golang check: "
if command -v go; then
    echo "    Version `go version` installed."
else
    echo "    golang not found, please install it and try again."
    exit 1
fi

# 拉取 git 仓库
echo ""
echo "👉 CiFa-install> 2/3 > Clone src..."

echo ""
echo "clone CiFa & CiFa-front from GitHub..."

mkdir CiFa
cd CiFa

mkdir dist
mkdir src
cd src

git clone https://github.com/cdfmlr/CiFa.git
git clone https://github.com/cdfmlr/CiFa-front.git

cd ..

# 编译
echo ""
echo "CiFa-install> 3/3 > Build src..."


# 编译后端
echo ""
echo "Build back-end..."

cd src/CiFa/main
go build main.go
echo "Done."
mv main ../../../dist/cifa
cd ../../..

# 编译前端
echo ""
echo "build front-end..."

cd src/Cifa-front
if command -v vue; then
    echo "vue cli installed. Skip."
else
    echo "vue cli missing. install it: npm install -g @vue/cli"
    npm install -g @vue/cli
fi
echo "npm install ant-design-vue"
npm install ant-design-vue
echo "npm run build"
npm run build
echo "build done."
mv dist ../../dist/static
cd ../../..

echo ""
echo "👉 CiFa-install>> Done."
echo "CiFa 安装在 ./CiFa/dist"
echo "开始使用: "
echo "    $ cd ./CiFa/dist  # 推荐进到目录再运行，或参考 ./CiFa/dist/cifa --help 手动配置 static 文件目录"
echo "    $ ./cifa"
echo "然后你可以在 http://localhost:9001 访问 CiFa 服务。"
echo "更多使用方法请参考：./cifa --help"
echo "---------------------------"
echo "Created by CDFMLR with ❤ . All rights reserved."
