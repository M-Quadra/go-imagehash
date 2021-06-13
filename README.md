# go-imagehash

[corona10/goimagehash](https://github.com/corona10/goimagehash)的大小设定不够灵活, 并且我更倾向于使用`golang.org/x/image`处理缩放

没啥用的修改, 妄图减少部分依赖

目前只稍微实现了`aHash`与`dHash`

# 李姐万岁

```
go get github.com/M-Quadra/go-imagehash
```

# 手到擒来

- aHash

```
aHashData, err := goimagehash.Sum.AHash(img)
```

自定义

```
aHash := goimagehash.NewAverageHash(8, 8, draw.CatmullRom)
aHashData, err := aHash.Sum(img)
```

- dHash

同理

# 疑难杂症

受限与不同实现对于缩放与灰度图处理不同, 最终结果会存在差异

有时颜色差异已经作为一种强特征, 但灰度图不适用

imageHash作用范围有限, 个人感觉已经够用, 若无必要后续只会随缘更新
