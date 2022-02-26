// Code generated by 'ccgo locale/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o locale/locale_linux_arm.go -pkgname locale', DO NOT EDIT.

package locale

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	LC_ADDRESS                = 9
	LC_ADDRESS_MASK           = 512
	LC_ALL                    = 6
	LC_ALL_MASK               = 8127
	LC_COLLATE                = 3
	LC_COLLATE_MASK           = 8
	LC_CTYPE                  = 0
	LC_CTYPE_MASK             = 1
	LC_IDENTIFICATION         = 12
	LC_IDENTIFICATION_MASK    = 4096
	LC_MEASUREMENT            = 11
	LC_MEASUREMENT_MASK       = 2048
	LC_MESSAGES               = 5
	LC_MESSAGES_MASK          = 32
	LC_MONETARY               = 4
	LC_MONETARY_MASK          = 16
	LC_NAME                   = 8
	LC_NAME_MASK              = 256
	LC_NUMERIC                = 1
	LC_NUMERIC_MASK           = 2
	LC_PAPER                  = 7
	LC_PAPER_MASK             = 128
	LC_TELEPHONE              = 10
	LC_TELEPHONE_MASK         = 1024
	LC_TIME                   = 2
	LC_TIME_MASK              = 4
	X_ATFILE_SOURCE           = 1
	X_BITS_LOCALE_H           = 1
	X_BITS_TYPES_LOCALE_T_H   = 1
	X_BITS_TYPES___LOCALE_T_H = 1
	X_DEFAULT_SOURCE          = 1
	X_FEATURES_H              = 1
	X_FILE_OFFSET_BITS        = 64
	X_LOCALE_H                = 1
	X_POSIX_C_SOURCE          = 200809
	X_POSIX_SOURCE            = 1
	X_STDC_PREDEF_H           = 1
	X_SYS_CDEFS_H             = 1
	Linux                     = 1
	Unix                      = 1
)

type Ptrdiff_t = int32 /* <builtin>:3:26 */

type Size_t = uint32 /* <builtin>:9:23 */

type Wchar_t = uint32 /* <builtin>:15:24 */

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Copyright (C) 1991-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

//	ISO C99 Standard: 7.11 Localization	<locale.h>

// Copyright (C) 1991-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// These are defined by the user (or the compiler)
//    to specify the desired environment:
//
//    __STRICT_ANSI__	ISO Standard C.
//    _ISOC99_SOURCE	Extensions to ISO C89 from ISO C99.
//    _ISOC11_SOURCE	Extensions to ISO C99 from ISO C11.
//    _ISOC2X_SOURCE	Extensions to ISO C99 from ISO C2X.
//    __STDC_WANT_LIB_EXT2__
// 			Extensions to ISO C99 from TR 27431-2:2010.
//    __STDC_WANT_IEC_60559_BFP_EXT__
// 			Extensions to ISO C11 from TS 18661-1:2014.
//    __STDC_WANT_IEC_60559_FUNCS_EXT__
// 			Extensions to ISO C11 from TS 18661-4:2015.
//    __STDC_WANT_IEC_60559_TYPES_EXT__
// 			Extensions to ISO C11 from TS 18661-3:2015.
//
//    _POSIX_SOURCE	IEEE Std 1003.1.
//    _POSIX_C_SOURCE	If ==1, like _POSIX_SOURCE; if >=2 add IEEE Std 1003.2;
// 			if >=199309L, add IEEE Std 1003.1b-1993;
// 			if >=199506L, add IEEE Std 1003.1c-1995;
// 			if >=200112L, all of IEEE 1003.1-2004
// 			if >=200809L, all of IEEE 1003.1-2008
//    _XOPEN_SOURCE	Includes POSIX and XPG things.  Set to 500 if
// 			Single Unix conformance is wanted, to 600 for the
// 			sixth revision, to 700 for the seventh revision.
//    _XOPEN_SOURCE_EXTENDED XPG things and X/Open Unix extensions.
//    _LARGEFILE_SOURCE	Some more functions for correct standard I/O.
//    _LARGEFILE64_SOURCE	Additional functionality from LFS for large files.
//    _FILE_OFFSET_BITS=N	Select default filesystem interface.
//    _ATFILE_SOURCE	Additional *at interfaces.
//    _GNU_SOURCE		All of the above, plus GNU extensions.
//    _DEFAULT_SOURCE	The default set of features (taking precedence over
// 			__STRICT_ANSI__).
//
//    _FORTIFY_SOURCE	Add security hardening to many library functions.
// 			Set to 1 or 2; 2 performs stricter checks than 1.
//
//    _REENTRANT, _THREAD_SAFE
// 			Obsolete; equivalent to _POSIX_C_SOURCE=199506L.
//
//    The `-ansi' switch to the GNU C compiler, and standards conformance
//    options such as `-std=c99', define __STRICT_ANSI__.  If none of
//    these are defined, or if _DEFAULT_SOURCE is defined, the default is
//    to have _POSIX_SOURCE set to one and _POSIX_C_SOURCE set to
//    200809L, as well as enabling miscellaneous functions from BSD and
//    SVID.  If more than one of these are defined, they accumulate.  For
//    example __STRICT_ANSI__, _POSIX_SOURCE and _POSIX_C_SOURCE together
//    give you ISO C, 1003.1, and 1003.2, but nothing else.
//
//    These are defined by this file and are used by the
//    header files to decide what to declare or define:
//
//    __GLIBC_USE (F)	Define things from feature set F.  This is defined
// 			to 1 or 0; the subsequent macros are either defined
// 			or undefined, and those tests should be moved to
// 			__GLIBC_USE.
//    __USE_ISOC11		Define ISO C11 things.
//    __USE_ISOC99		Define ISO C99 things.
//    __USE_ISOC95		Define ISO C90 AMD1 (C95) things.
//    __USE_ISOCXX11	Define ISO C++11 things.
//    __USE_POSIX		Define IEEE Std 1003.1 things.
//    __USE_POSIX2		Define IEEE Std 1003.2 things.
//    __USE_POSIX199309	Define IEEE Std 1003.1, and .1b things.
//    __USE_POSIX199506	Define IEEE Std 1003.1, .1b, .1c and .1i things.
//    __USE_XOPEN		Define XPG things.
//    __USE_XOPEN_EXTENDED	Define X/Open Unix things.
//    __USE_UNIX98		Define Single Unix V2 things.
//    __USE_XOPEN2K        Define XPG6 things.
//    __USE_XOPEN2KXSI     Define XPG6 XSI things.
//    __USE_XOPEN2K8       Define XPG7 things.
//    __USE_XOPEN2K8XSI    Define XPG7 XSI things.
//    __USE_LARGEFILE	Define correct standard I/O things.
//    __USE_LARGEFILE64	Define LFS things with separate names.
//    __USE_FILE_OFFSET64	Define 64bit interface as default.
//    __USE_MISC		Define things from 4.3BSD or System V Unix.
//    __USE_ATFILE		Define *at interfaces and AT_* constants for them.
//    __USE_GNU		Define GNU extensions.
//    __USE_FORTIFY_LEVEL	Additional security measures used, according to level.
//
//    The macros `__GNU_LIBRARY__', `__GLIBC__', and `__GLIBC_MINOR__' are
//    defined by this file unconditionally.  `__GNU_LIBRARY__' is provided
//    only for compatibility.  All new code should use the other symbols
//    to test for features.
//
//    All macros listed above as possibly being defined by this file are
//    explicitly undefined if they are not explicitly defined.
//    Feature-test macros that are not defined by the user or compiler
//    but are implied by the other feature-test macros defined (or by the
//    lack of any definitions) are defined by the file.
//
//    ISO C feature test macros depend on the definition of the macro
//    when an affected header is included, not when the first system
//    header is included, and so they are handled in
//    <bits/libc-header-start.h>, which does not have a multiple include
//    guard.  Feature test macros that can be handled from the first
//    system header included are handled here.

// Undefine everything, so we get a clean slate.

// Suppress kernel-name space pollution unless user expressedly asks
//    for it.

// Convenience macro to test the version of gcc.
//    Use like this:
//    #if __GNUC_PREREQ (2,8)
//    ... code requiring gcc 2.8 or later ...
//    #endif
//    Note: only works for GCC 2.0 and later, because __GNUC_MINOR__ was
//    added in 2.0.

// Similarly for clang.  Features added to GCC after version 4.2 may
//    or may not also be available in clang, and clang's definitions of
//    __GNUC(_MINOR)__ are fixed at 4 and 2 respectively.  Not all such
//    features can be queried via __has_extension/__has_feature.

// Whether to use feature set F.

// _BSD_SOURCE and _SVID_SOURCE are deprecated aliases for
//    _DEFAULT_SOURCE.  If _DEFAULT_SOURCE is present we do not
//    issue a warning; the expectation is that the source is being
//    transitioned to use the new macro.

// If _GNU_SOURCE was defined by the user, turn on all the other features.

// If nothing (other than _GNU_SOURCE and _DEFAULT_SOURCE) is defined,
//    define _DEFAULT_SOURCE.

// This is to enable the ISO C2X extension.

// This is to enable the ISO C11 extension.

// This is to enable the ISO C99 extension.

// This is to enable the ISO C90 Amendment 1:1995 extension.

// If none of the ANSI/POSIX macros are defined, or if _DEFAULT_SOURCE
//    is defined, use POSIX.1-2008 (or another version depending on
//    _XOPEN_SOURCE).

// Some C libraries once required _REENTRANT and/or _THREAD_SAFE to be
//    defined in all multithreaded code.  GNU libc has not required this
//    for many years.  We now treat them as compatibility synonyms for
//    _POSIX_C_SOURCE=199506L, which is the earliest level of POSIX with
//    comprehensive support for multithreaded code.  Using them never
//    lowers the selected level of POSIX conformance, only raises it.

// The function 'gets' existed in C89, but is impossible to use
//    safely.  It has been removed from ISO C11 and ISO C++14.  Note: for
//    compatibility with various implementations of <cstdio>, this test
//    must consider only the value of __cplusplus when compiling C++.

// GNU formerly extended the scanf functions with modified format
//    specifiers %as, %aS, and %a[...] that allocate a buffer for the
//    input using malloc.  This extension conflicts with ISO C99, which
//    defines %a as a standalone format specifier that reads a floating-
//    point number; moreover, POSIX.1-2008 provides the same feature
//    using the modifier letter 'm' instead (%ms, %mS, %m[...]).
//
//    We now follow C99 unless GNU extensions are active and the compiler
//    is specifically in C89 or C++98 mode (strict or not).  For
//    instance, with GCC, -std=gnu11 will have C99-compliant scanf with
//    or without -D_GNU_SOURCE, but -std=c89 -D_GNU_SOURCE will have the
//    old extension.

// Get definitions of __STDC_* predefined macros, if the compiler has
//    not preincluded this header automatically.
// Copyright (C) 1991-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// This macro indicates that the installed library is the GNU C Library.
//    For historic reasons the value now is 6 and this will stay from now
//    on.  The use of this variable is deprecated.  Use __GLIBC__ and
//    __GLIBC_MINOR__ now (see below) when you want to test for a specific
//    GNU C library version and use the values in <gnu/lib-names.h> to get
//    the sonames of the shared libraries.

// Major and minor version number of the GNU C library package.  Use
//    these macros to test for features in specific releases.

// This is here only because every header file already includes this one.
// Copyright (C) 1992-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// We are almost always included from features.h.

// The GNU libc does not support any K&R compilers or the traditional mode
//    of ISO C compilers anymore.  Check for some of the combinations not
//    anymore supported.

// Some user header file might have defined this before.

// All functions, except those with callbacks or those that
//    synchronize memory, are leaf functions.

// GCC can always grok prototypes.  For C++ programs we add throw()
//    to help it optimize the function calls.  But this works only with
//    gcc 2.8.x and egcs.  For gcc 3.2 and up we even mark C functions
//    as non-throwing using a function attribute since programs can use
//    the -fexceptions options for C code as well.

// Compilers that are not clang may object to
//        #if defined __clang__ && __has_extension(...)
//    even though they do not need to evaluate the right-hand side of the &&.

// These two macros are not used in glibc anymore.  They are kept here
//    only because some other projects expect the macros to be defined.

// For these things, GCC behaves the ANSI way normally,
//    and the non-ANSI way under -traditional.

// This is not a typedef so `const __ptr_t' does the right thing.

// C++ needs to know that types and declarations are C, not C++.

// Fortify support.

// Support for flexible arrays.
//    Headers that should use flexible arrays only if they're "real"
//    (e.g. only if they won't affect sizeof()) should test
//    #if __glibc_c99_flexarr_available.

// __asm__ ("xyz") is used throughout the headers to rename functions
//    at the assembly language level.  This is wrapped by the __REDIRECT
//    macro, in order to support compilers that can do this some other
//    way.  When compilers don't support asm-names at all, we have to do
//    preprocessor tricks instead (which don't have exactly the right
//    semantics, but it's the best we can do).
//
//    Example:
//    int __REDIRECT(setpgrp, (__pid_t pid, __pid_t pgrp), setpgid);

//
// #elif __SOME_OTHER_COMPILER__
//
// # define __REDIRECT(name, proto, alias) name proto; 	_Pragma("let " #name " = " #alias)

// GCC has various useful declarations that can be made with the
//    `__attribute__' syntax.  All of the ways we use this do fine if
//    they are omitted for compilers that don't understand it.

// At some point during the gcc 2.96 development the `malloc' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// Tell the compiler which arguments to an allocation function
//    indicate the size of the allocation.

// At some point during the gcc 2.96 development the `pure' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// This declaration tells the compiler that the value is constant.

// At some point during the gcc 3.1 development the `used' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// Since version 3.2, gcc allows marking deprecated functions.

// Since version 4.5, gcc also allows one to specify the message printed
//    when a deprecated function is used.  clang claims to be gcc 4.2, but
//    may also support this feature.

// At some point during the gcc 2.8 development the `format_arg' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.
//    If several `format_arg' attributes are given for the same function, in
//    gcc-3.0 and older, all but the last one are ignored.  In newer gccs,
//    all designated arguments are considered.

// At some point during the gcc 2.97 development the `strfmon' format
//    attribute for functions was introduced.  We don't want to use it
//    unconditionally (although this would be possible) since it
//    generates warnings.

// The nonull function attribute allows to mark pointer parameters which
//    must not be NULL.

// If fortification mode, we warn about unused results of certain
//    function calls which can lead to problems.

// Forces a function to be always inlined.
// The Linux kernel defines __always_inline in stddef.h (283d7573), and
//    it conflicts with this definition.  Therefore undefine it first to
//    allow either header to be included first.

// Associate error messages with the source location of the call site rather
//    than with the source location inside the function.

// GCC 4.3 and above with -std=c99 or -std=gnu99 implements ISO C99
//    inline semantics, unless -fgnu89-inline is used.  Using __GNUC_STDC_INLINE__
//    or __GNUC_GNU_INLINE is not a good enough check for gcc because gcc versions
//    older than 4.3 may define these macros and still not guarantee GNU inlining
//    semantics.
//
//    clang++ identifies itself as gcc-4.2, but has support for GNU inlining
//    semantics, that can be checked for by using the __GNUC_STDC_INLINE_ and
//    __GNUC_GNU_INLINE__ macro definitions.

// GCC 4.3 and above allow passing all anonymous arguments of an
//    __extern_always_inline function to some other vararg function.

// It is possible to compile containing GCC extensions even if GCC is
//    run in pedantic mode if the uses are carefully marked using the
//    `__extension__' keyword.  But this is not generally available before
//    version 2.8.

// __restrict is known in EGCS 1.2 and above.

// ISO C99 also allows to declare arrays as non-overlapping.  The syntax is
//      array_name[restrict]
//    GCC 3.1 supports this.

// Describes a char array whose address can safely be passed as the first
//    argument to strncpy and strncat, as the char array is not necessarily
//    a NUL-terminated string.

// Undefine (also defined in libc-symbols.h).
// Copies attributes from the declaration or type referenced by
//    the argument.

// Copyright (C) 1999-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Properties of long double type.
//    Copyright (C) 2016-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License  published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// This header is included by <sys/cdefs.h>.
//
//    If long double is ABI-compatible with double, it should define
//    __NO_LONG_DOUBLE_MATH to 1; otherwise, it should leave
//    __NO_LONG_DOUBLE_MATH undefined.
//
//    If this build of the GNU C Library supports both long double
//    ABI-compatible with double and some other long double format not
//    ABI-compatible with double, it should define
//    __LONG_DOUBLE_MATH_OPTIONAL to 1; otherwise, it should leave
//    __LONG_DOUBLE_MATH_OPTIONAL undefined.
//
//    If __NO_LONG_DOUBLE_MATH is already defined, this header must not
//    define anything; this is needed to work with the definition of
//    __NO_LONG_DOUBLE_MATH in nldbl-compat.h.

// In the default version of this header, long double is
//    ABI-compatible with double.

// __glibc_macro_warning (MESSAGE) issues warning MESSAGE.  This is
//    intended for use in preprocessor macros.
//
//    Note: MESSAGE must be a _single_ string; concatenation of string
//    literals is not supported.

// Generic selection (ISO C11) is a C-only feature, available in GCC
//    since version 4.9.  Previous versions do not provide generic
//    selection, even though they might set __STDC_VERSION__ to 201112L,
//    when in -std=c11 mode.  Thus, we must check for !defined __GNUC__
//    when testing __STDC_VERSION__ for generic selection support.
//    On the other hand, Clang also defines __GNUC__, so a clang-specific
//    check is required to enable the use of generic selection.

// If we don't have __REDIRECT, prototypes will be missing if
//    __USE_FILE_OFFSET64 but not __USE_LARGEFILE[64].

// Decide whether we can define 'extern inline' functions in headers.

// This is here only because every header file already includes this one.
//    Get the definitions of all the appropriate `__stub_FUNCTION' symbols.
//    <gnu/stubs.h> contains `#define __stub_FUNCTION' when FUNCTION is a stub
//    that will always return failure (and set errno to ENOSYS).
// This file is automatically generated.
//    This file selects the right generated file of `__stub_FUNCTION' macros
//    based on the architecture being compiled for.

// This file is automatically generated.
//    It defines a symbol `__stub_FUNCTION' for each function
//    in the C library which is a stub, meaning it will fail
//    every time called, usually setting errno to ENOSYS.

// Copyright (C) 1989-2020 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3, or (at your option)
// any later version.
//
// GCC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// ISO C Standard:  7.17  Common definitions  <stddef.h>

// Any one of these symbols __need_* means that GNU libc
//    wants us just to define one data type.  So don't define
//    the symbols that indicate this file's entire job has been done.

// This avoids lossage on SunOS but only if stdtypes.h comes first.
//    There's no way to win with the other order!  Sun lossage.

// Sequent's header files use _PTRDIFF_T_ in some conflicting way.
//    Just ignore it.

// On VxWorks, <type/vxTypesBase.h> may have defined macros like
//    _TYPE_size_t which will typedef size_t.  fixincludes patched the
//    vxTypesBase.h so that this macro is only defined if _GCC_SIZE_T is
//    not defined, and so that defining this macro defines _GCC_SIZE_T.
//    If we find that the macros are still defined at this point, we must
//    invoke them so that the type is defined as expected.

// In case nobody has defined these types, but we aren't running under
//    GCC 2.00, make sure that __PTRDIFF_TYPE__, __SIZE_TYPE__, and
//    __WCHAR_TYPE__ have reasonable values.  This can happen if the
//    parts of GCC is compiled by an older compiler, that actually
//    include gstddef.h, such as collect2.

// Signed type of difference of two pointers.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// Unsigned type of `sizeof' something.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// A null pointer constant.

// Definition of locale category symbol values.
//    Copyright (C) 2001-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// These are the possibilities for the first argument to setlocale.
//    The code assumes that the lowest LC_* symbol has the value zero.

// Structure giving information about numeric and monetary notation.
type Lconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    uint8
	Ffrac_digits        uint8
	Fp_cs_precedes      uint8
	Fp_sep_by_space     uint8
	Fn_cs_precedes      uint8
	Fn_sep_by_space     uint8
	Fp_sign_posn        uint8
	Fn_sign_posn        uint8
	Fint_p_cs_precedes  uint8
	Fint_p_sep_by_space uint8
	Fint_n_cs_precedes  uint8
	Fint_n_sep_by_space uint8
	Fint_p_sign_posn    uint8
	Fint_n_sign_posn    uint8
	F__ccgo_pad1        [2]byte
} /* locale.h:51:1 */

// POSIX.1-2008 extends the locale interface with functions for
//    explicit creation and manipulation of 'locale_t' objects
//    representing locale contexts, and a set of parallel
//    locale-sensitive text processing functions that take a locale_t
//    argument.  This enables applications to work with data from
//    multiple locales simultaneously and thread-safely.
// Definition of locale_t.
//    Copyright (C) 2017-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// Definition of struct __locale_struct and __locale_t.
//    Copyright (C) 1997-2020 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//    Contributed by Ulrich Drepper <drepper@cygnus.com>, 1997.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <https://www.gnu.org/licenses/>.

// POSIX.1-2008: the locale_t type, representing a locale context
//    (implementation-namespace version).  This type should be treated
//    as opaque by applications; some details are exposed for the sake of
//    efficiency in e.g. ctype functions.

type X__locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
} /* __locale_t.h:28:1 */

type X__locale_t = uintptr /* __locale_t.h:42:32 */

type Locale_t = X__locale_t /* locale_t.h:24:20 */

// This value can be passed to `uselocale' and may be returned by it.
//    Passing this value to any other function has undefined behavior.

var _ uint8 /* gen.c:2:13: */
