# simHtml

[![Build Status](https://travis-ci.com/cckuailong/simHtml.svg?branch=master)](https://travis-ci.com/cckuailong/simHtml)

[English Readme](https://github.com/cckuailong/py2so/blob/master/README.md) || 
[中文 Readme](https://github.com/cckuailong/py2so/blob/master/README_zh.md)

## 介绍

simHtml包提供了一些用于计算Web页面相似度的函数

## 安装

快速安装:

    go get -u github.com/cckuailong/simHtml

## 原理

### 网页结构相似度

使用序列比较方法（最长公共子序列）来计算dom树的相似度。

### 元素类型相似度

计算class 和 style的相似度。

### 整合 网页结构相似度 和 元素类型相似度

整合算法：

    k * structural_similarity(document_1, document_2) + (1 - k) * style_similarity(document_1, document_2)

相似度取值在0-1之间

### k取值建议

使用 `k=0.3` 可以获得更好的结果。 元素类型相似度 包含的信息更多，更精确。

## 函数

- GetSimFromFile(file1, file2 string) float64

```
    In [1]: 1.html's content is
     '''
     <html>
        <h1 class="title">First Document</h1>
        <ul class="menu">
            <li class="active">Documents</li>
            <li>Extra</li>
        </ul>
     </html>
    '''

    In [2]: 2.html's content is
    '''
    <html>
        <h1 class="title">Second document Document</h1>
        <ul class="menu">
            <li class="active">Extra Documents</li>
        </ul>
    </html>
    '''

    In [3] import "github.com/cckuailong/simHtml/simHtml"

    In [4]: simHtml.GetSimRate("./1.html", "./2.html")
    Out[4]: 0.9727272727272727
```

- GetSimFromStr(str1, str2 string) float64

- GetSimFromUrl(url1, url2 string) float64

### 参考

- html-similarity <https://github.com/matiskay/html-similarity>