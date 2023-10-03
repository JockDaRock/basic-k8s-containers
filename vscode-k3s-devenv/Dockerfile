FROM codercom/code-server

# For Now I am switching the USER back to root to make running docker in docker easier
USER root

# Installing Ubuntu Dependancies
# RUN apt-get update && apt-get upgrade -y python3 && apt-get install -y python3-pip git

# Install code language dependancies
# RUN pip3 install requests bottle

# Docker engine install
RUN apt update && apt install -y \
    apt-transport-https \
    gnupg \
    ca-certificates \
    curl \
    wget \
    software-properties-common \
    && install -m 0755 -d /etc/apt/keyrings
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg \
    && chmod a+r /etc/apt/keyrings/docker.gpg && echo \
    "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
    tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && apt-get update && apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# ioxclient install
# RUN curl -O https://pubhub.devnetcloud.com/media/iox/docs/artifacts/ioxclient/ioxclient-v1.16.0.0/ioxclient_1.16.0.0_linux_amd64.tar.gz \
#     && tar xvzf ioxclient_1.16.0.0_linux_amd64.tar.gz \
#     && chmod +x ioxclient_1.16.0.0_linux_amd64/ioxclient \
#     && cp ioxclient_1.16.0.0_linux_amd64/ioxclient /usr/local/bin/ioxclient

# Powershell install
# RUN wget -q https://packages.microsoft.com/config/ubuntu/20.04/packages-microsoft-prod.deb \
#     && dpkg -i packages-microsoft-prod.deb \
#     && apt update && apt install -y powershell


# VS Code Extensions Install for code-server - Powershell Terminal Extension
#RUN apt update && apt install -y libarchive-tools
#RUN apt update && apt install -y bsdutils
#RUN mkdir -p /root/.local/share/code-server/extensions
#RUN curl -JL https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ms-vscode/vsextensions/PowerShell/2021.12.0/vspackage | bsdtar -xvf - extension \
#    && mv extension /root/.local/share/code-server/extensions/ms-vscode.powershell

#RUN code --install-extension powershell-2021.12.0.vsix

# Installing k3s without autostart
# RUN wget -O /usr/bin/k3s https://github.com/k3s-io/k3s/releases/download/v1.28.2%2Bk3s1/k3s
# RUN chmod +x /usr/bin/k3s
# RUN curl -sfL https://get.k3s.io | INSTALL_K3S_SKIP_START=true INSTALL_K3S_SKIP_ENABLE=true sh -s - --docker
RUN curl -sfL https://get.k3s.io | INSTALL_K3S_SKIP_START=true INSTALL_K3S_SKIP_ENABLE=true sh -

# Installing helm
RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 \
    && chmod 700 get_helm.sh && ./get_helm.sh


# Coping start script
COPY start_codedock.sh /usr/local/bin/start_codedock.sh

# Copy vscode settings
COPY settings.json /root/.local/share/code-server/User/settings.json
COPY coder.json /root/.local/share/code-server/coder.json
COPY config.yaml /root/.config/code-server/config.yaml

RUN chmod +x /usr/local/bin/start_codedock.sh

ENTRYPOINT ["/usr/bin/dumb-init", "start_codedock.sh"]