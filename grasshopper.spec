Name:           grasshopper
Version:        0.0.4
Release:        1%{?dist}
Summary:        This will make a Nulecule GO!

License:        LGPLv3+
URL:            https://github.com/goern/grasshopper
Source0:        https://github.com/goern/grasshopper-0.0.4.tar.gz

%description
This will make a Nulecule GO!

%prep
%setup -q -n %{name}-%{version}

mkdir -p $RPM_BUILD_ROOT/%{_bindir}

%build
make
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
alternatives --install %{_bindir}/grasshopper grasshopper %{_bindir}/grasshopper-{goversion} %{alternatives_priority}

%preun
alternatives --remove grasshopper %{_bindir}/grasshopper-%{goversion}

%changelog
* Tue Nov  3 2015 Christoph Görn
- initial RPMification
