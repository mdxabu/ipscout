## Vulnerabilities Found
=== Symbol Results ===

Vulnerability #1: GO-2025-3447
    Timing sidechannel for P-256 on ppc64le in crypto/internal/nistec
  More info: https://pkg.go.dev/vuln/GO-2025-3447
  Standard library
    Found in: crypto/internal/nistec@go1.23.2
    Fixed in: crypto/internal/nistec@go1.23.6
    Platforms: ppc64le
    Example traces found:
      #1: cmd/find.go:8:2: cmd.init calls ip2location.init, which eventually calls nistec.P256Point.SetBytes

Vulnerability #2: GO-2025-3373
    Usage of IPv6 zone IDs can bypass URI name constraints in crypto/x509
  More info: https://pkg.go.dev/vuln/GO-2025-3373
  Standard library
    Found in: crypto/x509@go1.23.2
    Fixed in: crypto/x509@go1.23.5
    Example traces found:
      #1: cmd/find.go:8:2: cmd.init calls ip2location.init, which eventually calls x509.CertPool.AppendCertsFromPEM
      #2: cmd/root.go:19:24: cmd.Execute calls cobra.Command.Execute, which eventually calls x509.HostnameError.Error
      #3: cmd/find.go:8:2: cmd.init calls ip2location.init, which eventually calls x509.ParseCertificate

Your code is affected by 2 vulnerabilities from the Go standard library.
This scan also found 1 vulnerability in packages you import and 0
vulnerabilities in modules you require, but your code doesn't appear to call
these vulnerabilities.
Use '-show verbose' for more details.
