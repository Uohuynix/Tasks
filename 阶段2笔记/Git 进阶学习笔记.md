# Git 进阶学习笔记

## 1. 分支管理

### 1.1 创建分支

```
bash# 创建新分支并切换到该分支
git checkout -b <branch_name>

# 或者分步执行
git branch <branch_name>
git checkout <branch_name>
```

### 1.2 查看分支

```
bash# 查看本地分支
git branch

# 查看远程分支
git branch -r

# 查看所有分支（本地和远程）
git branch -a
```

### 1.3 切换分支

```
bashgit checkout <branch_name>
```

### 1.4 合并分支

```
bash# 切换到目标分支
git checkout <target_branch>

# 合并指定分支到当前分支
git merge <source_branch>
```

### 1.5 解决冲突

在合并分支时，可能会发生冲突。解决冲突的步骤：

1. 手动解决冲突文件的内容。
2. 使用 `git add` 将解决后的文件标记为已解决。
3. 使用 `git merge --continue` 继续合并过程。

### 1.6 删除分支

```
bash# 删除本地分支
git branch -d <branch_name>

# 强制删除分支（未合并的分支）
git branch -D <branch_name>
```

## 2. 高级版本控制

### 2.1 交互式暂存

```
bash# 选择要暂存的文件
git add -i
```

### 2.2 暂存部分文件内容

```
bash# 部分文件暂存
git add -p
```

### 2.3 修改最后一次提交

```
bash# 修改最后一次提交的信息
git commit --amend
```

### 2.4 重置提交历史

```
bash# 软重置，保留工作目录和暂存区，将 HEAD 指向新的提交
git reset --soft <commit_sha>

# 混合重置，保留工作目录，清空暂存区，将 HEAD 指向新的提交
git reset --mixed <commit_sha>

# 硬重置，工作目录和暂存区都被清空，将 HEAD 指向新的提交
git reset --hard <commit_sha>
```

### 2.5 Cherry-pick

```
bash# 将指定提交应用到当前分支
git cherry-pick <commit_sha>
```

## 3. 远程操作

### 3.1 添加远程仓库

```
bashgit remote add <remote_name> <remote_url>
```

### 3.2 查看远程仓库

```
bashgit remote -v
```

### 3.3 拉取远程分支

```
bash# 拉取远程分支到本地，并自动创建对应的本地分支
git checkout -b <local_branch_name> origin/<remote_branch_name>
```

### 3.4 推送分支

```
bash# 推送本地分支到远程仓库
git push origin <branch_name>
```

### 3.5 删除远程分支

```
bashgit push origin --delete <branch_name>
```

## 4. Git 钩子

### 4.1 客户端钩子

- `pre-commit`: 在执行提交前运行，用于检查即将提交的内容。
- `pre-receive`: 在服务器接收提交前运行，用于检查即将到来的推送。
- `post-commit`: 在提交执行后运行，用于执行额外的操作。

### 4.2 服务端钩子

- `pre-receive`: 在服务器接收提交前运行，用于检查即将到来的推送。
- `post-receive`: 在服务器接收提交后运行，用于执行额外的操作。

## 5. Git 工作流

### 5.1 Git Flow

Git Flow 是一种流行的分支管理模型，定义了一套基于分支的协作规范，包括 `feature`, `release`, `hotfix` 等分支。

### 5.2 GitHub Flow

GitHub Flow 是一种更简化的工作流，主要包括从 `main` 分支创建 `feature` 分支，开发完毕后合并回 `main`。

## 总结

Git 是开发中不可或缺的版本控制工具，通过学习和掌握进阶的 Git 技巧，能够更高效地管理代码版本、处理分支、解决冲突，并与团队协作。熟练使用 Git 工具和理解相关概念，将有助于提高代码管理和协作的效率。



## 引用

CSDN

必应