# Git Comment 书写规范

基本格式示例

```
[+] feat: <log/stack> 添加 XXXX 功能
```

上述示例有几个段落, 每个段落之间都有空格分割, 下面说明各段落作用:

| 示例 | 作用 | 详细说明 |
|:---:|:---:|:---|
| `[+]` | 增删改符号 | `[+]` 添加；`[-]` 删除；`[*]` 更新 |
| `feat:` | 变动类别 | |
| `<log/stack>` | 子模块 | 如果是根目录则不需要这个字段 |
| 剩余 | log 正文 | |

变动类别说明:

| 类别 | 说明 |
|:---:|:---|
| feat | 功能性变更, 可以是产品需求也可以是研发需求。包括新增、修改、删除 |
| perf | 优化相关，比如提升性能 (performance) |
| fix | 表示修复 bug 或调整不合理的逻辑 |
| docs | 新增或修改文档和注释 |
| test | 测试用例，包括单元测试、集成测试等 |
| revert | 回滚代码 |
| refactor | 代码重构，没有加新功能或者修复 bug |
| style | 仅仅修改了空格、格式缩进、逗号等等，不改变代码逻辑 |
| chore | 上述分类未能覆盖的杂项变更 |
