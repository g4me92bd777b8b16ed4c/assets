#!/bin/bash
find assets -depth -exec rename 's/(.*)\/([^\/]*)/$1\/\L$2/' {} \;
