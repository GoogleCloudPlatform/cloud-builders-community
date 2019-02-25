#!/usr/bin/env sh

### Path Ruby RBENV / RVM ###
export RBENV_ROOT="$HOME/.rbenv"
export RVM_ROOT="/usr/local/rvm"

### rbenv (Ruby) default ###
if [ -d "$RBENV_ROOT" ]
then
  export PATH="$RBENV_ROOT/bin:${PATH}"
  eval "$(rbenv init -)"
  export PATH="$RBENV_ROOT/plugins/ruby-build/bin:$PATH"
else
  ### rvm (Ruby) - alternative ###
  if [ -d "$RVM_ROOT" ]
  then
    export PATH="$PATH:$RVM_ROOT/bin"
    source $RVM_ROOT/scripts/rvm

    # Set PATH alternatives using this:
    [[ -s "$RVM_ROOT/scripts/rvm"  ]] && source "$RVM_ROOT/scripts/rvm"
  fi
fi
