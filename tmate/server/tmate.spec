Name:           tmate-server
Version:        1.8
Release:        1%{?dist}

Summary:        Instant terminal sharing
License:        MIT
Url:            http://tmate.io

BuildRequires:  autoconf
BuildRequires:  libtool
BuildRequires:  pkgconfig
BuildRequires:  ruby
BuildRequires:  libevent-devel
BuildRequires:  openssl-devel
BuildRequires:  ncurses-devel
BuildRequires:  zlib-devel
BuildRequires:  libssh-devel >= 0.6.0
BuildRequires:  msgpack-devel >= 0.5.8

Source0:        https://github.com/nviennot/tmate-slave/archive/1.8.zip

%description
Tmate is a fork of tmux providing an instant pairing solution.

%prep
%setup -q -n %{name}-%{version}

%build
./autogen.sh
%configure
make %{?_smp_mflags}

%install
make DESTDIR=%{buildroot} install

%files
%defattr(-,root,root)
%doc CHANGES FAQ README-tmux README.md
%{_bindir}/tmate
%{_mandir}/man1/tmate.1*

%changelog

* Sat Jan 24 2015 - Kevin Mulvey <kmulvey@linux.com> - 1.8-1
- The big bang.
