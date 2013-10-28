clear
echo " "
echo " "
echo "                                  - Boxen -"
echo "                       QUIT WORRYING ABOUT YOUR TOOLS."  
echo "           Automate the pain out of your development environment."  
echo "  Boxen installs your dependencies so you can focus on getting things done."
echo " "
echo "  boxen-portal includes adium, caffeine, google-chrome, homebrew, java," 
echo "  netbeans, phpstorm, rbenv, ruby, subversion 1.7 vagrant 1.2,"
echo "  vagrant-vbguest and virtualbox"
echo " "
echo " "
mkdir ~/boxen
git clone -b 2.0 https://github.com/healeyious/boxen-portal ~/boxen/repo
cd ~/boxen/repo
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
ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"
echo "... rbenv and ruby ..."
brew install rbenv
brew install ruby
