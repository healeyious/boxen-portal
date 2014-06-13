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
echo "fix ruby"
curl https://gist.githubusercontent.com/Paulche/9713531/raw/1e57fbb440d36ca5607d1739cc6151f373b234b6/gistfile1.txt | sudo patch /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/2.0.0/universal-darwin13/rbconfig.rb
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
