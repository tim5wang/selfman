团队遇到一个很顽固的问题：
> Cleaning up project directory and file based variables
30:00
ERROR: Job failed: exit status 1

问题排查：

1. 怀疑是没权限，导致无法删除相关目录，实际在ci配置里写删除文件命令，是可以删除掉的，并且仍然会报上述错误

2. 怀疑是版本不兼容：官方文档有明确说明版本要适配，需要大版本一致
    - gitlab版本: 13.6.7-ee
    - gitlab-runner版本： 14.3.0

    结合着两个文档找到合适的镜像：
    - https://gitlab.com/gitlab-org/gitlab-runner/-/tags?page=2&sort=updated_desc
    - https://s3.amazonaws.com/gitlab-runner-downloads/13.6.7/binaries/gitlab-runner-linux-386

    https://s3.amazonaws.com/gitlab-runner-downloads/v13.12.0/binaries/gitlab-runner-linux-386

    经过实践，13.6.0 版本的gitlab-runner问题依旧
    
    所以尝试 13.12.0 版本的



