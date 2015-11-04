Name:           grasshopper
Version:        0.0.6
Release:        1%{?dist}
Summary:        This will make a Nulecule GO!

License:        LGPLv3+
URL:            https://github.com/goern/grasshopper
Source0:        grasshopper-0.0.5.tar.gz

BuildRequires:	golang-bin
Requires:       golang

%ifarch x86_64
    %global GOARCH amd64
%endif

%description
This will make a Nulecule GO!

%prep
%setup -q -n %{name}-%{version}

mkdir -p $RPM_BUILD_ROOT/%{_bindir}

%build
GOROOT="$(pwd)"
GOSRC="$(GOROOT)"
GOPATH="$(GOROOT)"
GOBIN="$GOROOT/bin"
GOOS=linux
GOARCH="%{GOARCH}"
export GOSRC GOROOT GOOS GOBIN

LC_ALL=C PATH="$PATH:$GOBIN" make
cp grasshopper-%{version} $RPM_BUILD_ROOT/%{_bindir}/grasshopper-%{version}

%clean
[ "$RPM_BUILD_ROOT" != "/" ] && rm -rf $RPM_BUILD_ROOT

%files
%defattr(0644,root,root,0755)

%attr(0755,-,-) %{_bindir}/grasshopper-%{version}

%doc AUTHORS
%doc LICENSE
%doc README.md

%post
alternatives --install %{_bindir}/grasshopper grasshopper %{_bindir}/grasshopper-{version} %{alternatives_priority}

%preun
alternatives --remove grasshopper %{_bindir}/grasshopper-%{version}

%changelog
* Wed Nov 04 2015 Christoph Görn <goern@redhat.com> 0.0.6-1
- 

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com>
- move to rel-eng/ (goern@redhat.com)
- add man page (stub) (goern@redhat.com)
- remove the .srpm (goern@redhat.com)

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com>
- move to rel-eng/ (goern@redhat.com)
- add man page (stub) (goern@redhat.com)
- remove the .srpm (goern@redhat.com)

* Tue Nov 03 2015 Christoph Görn <goern@redhat.com> 0.0.5-1
- new package built with tito

* Tue Nov 03 2015 Christoph Görn
- initial RPMification
