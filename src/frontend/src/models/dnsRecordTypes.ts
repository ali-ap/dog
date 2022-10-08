enum DNSRecordType {
    A = 1,
    NS = 2,
    CNAME = 5,
    SOA = 6,
    PTR = 12,
    MX = 15,
    TXT = 16,
    AAAA = 28,
    SRV = 33,
    DS = 43,
    DNSKEY = 48,
    TLSA = 52,
    TSIG = 250,
    ANY = 255,
    CAA = 257
}

export default DNSRecordType;