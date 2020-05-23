FROM hub.c.163.com/netease_comb/centos:7
MAINTAINER netease

ADD TDengine-client-1.6.5.9-Linux-x64.tar.gz /usr/local/

RUN ln -s /usr/local/TDengine-client/driver/libtaos.so.1.6.5.9 /usr/lib/libtaos.so

# 更新yum源
RUN yum makecache fast && yum -y update glibc

# 安装常用软件
RUN yum install -y openssh-server vim tar wget curl rsync bzip2 iptables tcpdump less telnet net-tools lsof
# 初始化ssh登陆
RUN ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N ''
RUN ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key -N ''
RUN ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N ''
RUN ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N ''
RUN echo "RSAAuthentication yes" >> /etc/ssh/sshd_config
RUN echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config

RUN yum clean all

# 启动sshd服务并且暴露22端口
RUN mkdir /var/run/sshd

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]