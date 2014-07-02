clear
echo "首先，让我们"
java -version
xcode-select --install
rm -Rf ~/boxen
mkdir ~/boxen
git clone https://github.com/healeyious/boxen-portal ~/boxen/repo
cd ~/boxen/repo
echo "... hosts ..."
sudo cp ~/boxen/repo/int-hosts /etc/hosts
echo "... ruby and json ..."
ARCHFLAGS=-Wno-error=unused-command-line-argument-hard-error-in-future ./script/boxen --no-fde
touch ~/.bashrc
echo '[ -f /opt/boxen/env.sh ] && source /opt/boxen/env.sh' | cat > ~/.bashrc
chown ${USER}:staff ~/.bashrc && chmod 755 ~/.bashrc
echo "... vagrant-vbguest ..."
vagrant plugin install vagrant-vbguest
echo "... homebrew ..."
ruby -e "$(curl -fsSL https://raw.github.com/Homebrew/homebrew/go/install)"
echo "... nodejs ..."
brew tap homebrew/versions
brew install node
echo "... rbenv and ruby ..."
brew install rbenv
brew install ruby
echo "... subversion ..."
brew tap homebrew/versions
brew install subversion
sudo rm -rf /usr/bin/svn*
sudo ln -s /usr/local/Cellar/subversion/1.8.9/bin/svn* /usr/bin/
subversion --version
