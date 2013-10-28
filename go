clear
echo d8888b.  .d88b.  db    db d88888b d8b   db 
echo 88  `8D .8P  Y8. `8b  d8' 88'     888o  88 
echo 88oooY' 88    88  `8bd8'  88ooooo 88V8o 88 
echo 88~~~b. 88    88  .dPYb.  88~~~~~ 88 V8o88 
echo 88   8D `8b  d8' .8P  Y8. 88.     88  V888 
echo Y8888P'  `Y88P'  YP    YP Y88888P VP   V8P 
echo .
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
