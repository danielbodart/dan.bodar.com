---
title: "Bash function to switch Java versions"
date: 2014-03-27T10:39:13Z
slug: "bash-function-to-switch-java-versions"
categories:
  - java
  - memory
---

```

setjava() {

	if [ "$1" = "-q" ]; then

                local quiet=true

                shift

        fi

	local jdk=~/Applications/Java/jdk1.$1

        if [ ! -d "${jdk}" ]; then

                echo Jdk not found: ${jdk}

                return 1

        fi

	export JAVA_HOME=${jdk}

        export PATH=${JAVA_HOME}/bin:${PATH}

        if [ -z "${quiet}" ]; then

                java -version

        fi

}

export -f setjava

```

I have symlinks for all major versions of Java so that in IntelliJ and the command line I can upgrade minor Java versions just by changing the symlink:

```

$ ls -la ~/Applications/Java/

total 20

drwxrwxr-x. 5 dan dan 4096 Mar 27 10:12 .

drwxrwxr-x. 8 dan dan 4096 Mar 27 10:13 ..

lrwxrwxrwx. 1 dan dan   11 Mar 27 10:12 jdk1.6 -> jdk1.6.0_45

drwxr-xr-x. 8 dan dan 4096 Mar 26  2013 jdk1.6.0_45

lrwxrwxrwx. 1 dan dan   11 Mar 18 13:55 jdk1.7 -> jdk1.7.0_51

drwxr-xr-x. 8 dan dan 4096 Dec 19 03:24 jdk1.7.0_51

lrwxrwxrwx. 1 dan dan    8 Mar 27 10:12 jdk1.8 -> jdk1.8.0

drwxr-xr-x. 8 dan dan 4096 Mar  4 11:18 jdk1.8.0

```

Then I can switch version by either just saying the major version:

```

$ setjava 8

java version "1.8.0"

Java(TM) SE Runtime Environment (build 1.8.0-b132)

Java HotSpot(TM) 64-Bit Server VM (build 25.0-b70, mixed mode)

```

or I can be specific about minor version:

```

$ setjava 6.0_45

java version "1.6.0_45"

Java(TM) SE Runtime Environment (build 1.6.0_45-b06)

Java HotSpot(TM) 64-Bit Server VM (build 20.45-b01, mixed mode)

```

I've also started adding the following to my build and execution scripts so they can specify the Java version themselves:

```

type -t setjava > /dev/null && setjava -q 8 || if [ -n "${JAVA_HOME}" ]; then PATH=${JAVA_HOME}/bin:${PATH}; fi

```

Currently this is designed to be optional and will fallback to JAVA\_HOME etc.