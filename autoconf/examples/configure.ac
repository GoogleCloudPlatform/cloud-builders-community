AC_INIT([hello], [1.0], [your-email@example.com])
AM_INIT_AUTOMAKE([-Wall -Werror foreign])
AC_CONFIG_SRCDIR([hello.c])
AC_CONFIG_HEADERS([config.h])

# Checks for programs.
AC_PROG_CC

# Optional feature (example)
AC_ARG_ENABLE([greeting],
              [AS_HELP_STRING([--enable-greeting],
                              [enable a more verbose greeting])],
              [enable_greeting=$enableval],
              [enable_greeting=no])

if test "x$enable_greeting" = "xyes"; then
  AC_DEFINE([ENABLE_GREETING], [1], [Define to enable the verbose greeting.])
fi

# Output files to generate
AC_CONFIG_FILES([Makefile])

# Output summary of configuration
AC_OUTPUT
