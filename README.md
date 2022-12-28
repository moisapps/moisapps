# moisapps
Moisapps is an automation for application creation and release

```shell
sudo dpkg -r moisapps && \
rm -rf build/package/moisapps.deb build/package/moisapps/usr/local/bin/moisapps && \
go build -o build/package/moisapps/usr/local/bin/moisapps && \
cd build/package && \
dpkg-deb --build --root-owner-group moisapps && \
cd ../../ && \
sudo dpkg -i build/package/moisapps.deb && \
echo 'source <(moisapps completion bash)' >> ~/.bashrc && \
sudo bash -c 'moisapps completion bash > /etc/bash_completion.d/moisapps' && \
echo 'source /usr/share/bash-completion/bash_completion' >>~/.bashrc
```