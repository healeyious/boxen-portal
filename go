clear
echo "."
echo "."
echo "      BBBBBBB      -Portal Edition-"
echo "      B     B"
echo "      B     B  OOOOO  X   X  EEEEE  N    N"
echo "      BBBBBBB  O   O   X X   E      N N  N"
echo "      B     B  O   O    X    EEEEE  N  N N"  
echo "      B     B  O   O   X X   E      N   NN"  
echo "      BBBBBBB  OOOOO  X   X  EEEEE  N    N"
echo "."
echo "        QUIT WORRYING ABOUT YOUR TOOLS."
echo "      Automate the pain out of your development environment."
echo "      Boxen installs your dependencies so you can focus on getting things done."
echo "."
echo "."
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
