clear
echo " "
echo " "
echo "首先，让我们"
echo " "
echo " "
java -version
xcode-select --install
rm -Rf ~/boxen
mkdir ~/boxen
git clone -b 2.0 https://github.com/healeyious/boxen-portal ~/boxen/repo
cd ~/boxen/repo
echo "... ruby and json ..."
ARCHFLAGS=-Wno-error=unused-command-line-argument-hard-error-in-future ./script/boxen --no-fde
touch ~/.bashrc
echo '[ -f /opt/boxen/env.sh ] && source /opt/boxen/env.sh' | cat > ~/.bashrc
chown ${USER}:staff ~/.bashrc && chmod 755 ~/.bashrc
echo "... subversion ..."
brew tap homebrew/versions
brew install subversion18
echo "... vagrant-vbguest ..."
vagrant plugin install vagrant-vbguest
echo "... nodejs ..."
brew tap homebrew/versions
brew install node
echo "... homebrew ..."
ruby -e "$(curl -fsSL https://raw.github.com/Homebrew/homebrew/go/install)"
echo "... rbenv and ruby ..."
brew install rbenv
brew install ruby
