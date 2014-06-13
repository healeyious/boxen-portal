clear
echo " "
echo " "
echo "                          Hi,"
echo " "
echo "                      Let's begin"
echo "                Lassen Sie uns beginnen"
echo "                      Commençons"
echo "                     首先，让我们"
echo "                     始めましょう"
echo " "
echo " "
rm -Rf ~/boxen
mkdir ~/boxen
git clone -b 2.0 https://github.com/healeyious/boxen-portal ~/boxen/repo
cd ~/boxen/repo
echo "... ruby and json ..."
script/boxen --no-fde
touch ~/.bashrc
echo '[ -f /opt/boxen/env.sh ] && source /opt/boxen/env.sh' | cat > ~/.bashrc
chown ${USER}:staff ~/.bashrc && chmod 755 ~/.bashrc
echo "... subversion 1.7 ..."
brew tap homebrew/versions
brew install subversion17
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
