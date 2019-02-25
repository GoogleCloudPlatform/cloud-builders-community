#!/usr/bin/env sh

### Path RVM ###
export RVM_ROOT="/usr/local/rvm"
export INSTALL_VERSION=${RUBY_VERSION:-2.6.1}
export RUBY_PATH=$RVM_ROOT/rubies/ruby-$INSTALL_VERSION

curl -sSL https://rvm.io/mpapis.asc | gpg2 --import -
curl -sSL https://rvm.io/pkuczynski.asc | gpg2 --import -

curl -sSL https://raw.githubusercontent.com/wayneeseguin/rvm/master/binscripts/rvm-installer | bash -s stable
$RVM_ROOT/bin/rvm get head --auto-dotfiles

#-----------------------------------------------------------------------------
# Install Ruby with RVM
#-----------------------------------------------------------------------------
$RVM_ROOT/bin/rvm install $INSTALL_VERSION
$RVM_ROOT/bin/rvm use $INSTALL_VERSION --default

ln -sf $RVM_ROOT/bin/rvm /usr/local/bin/rvm
ln -sf $RUBY_PATH/bin/ruby /usr/local/bin/ruby
ln -sf $RUBY_PATH/bin/gem /usr/local/bin/gem

#gem install bundler
