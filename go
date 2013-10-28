echo "... things ..."
mkdir ~/boxen
git clone -b 2.0 https://github.com/healeyious/boxen-portal ~/boxen/repo
cd ~/boxen/repo
script/boxen --no-fde
touch ~/.bashrc
echo '[ -f /opt/boxen/env.sh ] && source /opt/boxen/env.sh' | cat > ~/.bashrc
chown ${USER}:staff ~/.bashrc && chmod 755 ~/.bashrc
echo "... subversion17 ..."
brew tap homebrew/versions
brew install subversion17
echo "... vagrant-vbguest ..."
vagrant plugin install vagrant-vbguest
echo "... nodejs ..."
brew tap homebrew/versions
brew install node
echo "... homebrew ..."
ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"
echo "Installling rbenv and ruby (requires homebrew)..."
brew install rbenv
brew install ruby
