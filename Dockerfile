# Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/Daz-3ux/dBlog.

FROM archlinux/archlinux:base-devel AS base

LABEL maintainer="<daz-3ux@proton.me>"

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    DB_HOST=<127.0.0.1> \
    DB_PORT=<3306> \
    DB_USER=<root> \
    DB_PASSWORD=<passwd> \
    DB_NAME=<dazBlog>
ENV TZ Asia/Shanghai

RUN echo "Server = https://mirrors.bfsu.edu.cn/archlinux/\$repo/os/\$arch" | tee /etc/pacman.d/mirrorlist && \
    pacman -Syyu --noconfirm && \
    pacman -S --noconfirm go && \
    pacman -Scc --noconfirm

FROM base AS build

WORKDIR /dBlog

COPY . .

RUN make tools.verify && \
    make ca && \
    make && \
    chmod +r /dBlog/internal/resource/404.html && \
    chmod +x /dBlog/_output/platforms/linux/amd64/dBlog

FROM ubuntu:devel AS final

WORKDIR /dBlog

RUN  mkdir -p /var/log/dazblog &&\
     mkdir -p /etc/ssl/certs/ &&\
     touch /var/log/dazblog/dazblog.log && \
     apt-get update  && \
     apt-get install curl -y


COPY --from=build /dBlog/_output/platforms/linux/amd64/dBlog /dBlog/_output/platforms/linux/amd64/dBlog
COPY --from=build /dBlog/_output/cert/ /dBlog/_output/cert
COPY --from=build /dBlog/internal/resource/404.html /dBlog/internal/resource/404.html
COPY --from=build /dBlog/configs/dazBlog.yaml /dBlog/configs/dazBlog.yaml

HEALTHCHECK CMD curl --fail http://localhost:8081/healthz || exit 1

ENTRYPOINT ["_output/platforms/linux/amd64/dBlog"]
CMD ["-c", "configs/dazBlog.yaml"]
