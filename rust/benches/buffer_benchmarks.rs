use criterion::{criterion_group, criterion_main, Criterion};

use buffer_benchmarks::{flatbuffers, prost, protobuf};

fn bench_protobuf(c: &mut Criterion) {
    let m = protobuf::prepare_msg();
    c.bench_function("encode_protobuf", |b| b.iter(|| protobuf::encode(&m)));
    let buf = protobuf::encode(&m);
    println!("Wire format size (bytes) = {}", buf.len());
    c.bench_function("decode_protobuf", |b| b.iter(|| protobuf::decode(&buf)));
}

fn bench_prost(c: &mut Criterion) {
    let m = prost::prepare_msg();
    c.bench_function("encode_prost", |b| b.iter(|| prost::encode(&m)));
    let buf = prost::encode(&m);
    println!("Wire format size (bytes) = {}", buf.len());
    c.bench_function("decode_prost", |b| b.iter(|| prost::decode(&buf)));
}

fn bench_flatbuffers(c: &mut Criterion) {
    let m = flatbuffers::prepare_msg();
    c.bench_function("encode_flatbuffers", |b| b.iter(|| flatbuffers::encode(&m)));
    let buf = flatbuffers::encode(&m);
    println!("Wire format size (bytes) = {}", buf.finished_data().len());
    c.bench_function("decode_flatbuffers", |b| {
        b.iter(|| flatbuffers::decode(buf.finished_data()))
    });
}

criterion_group!(benches, bench_protobuf, bench_prost, bench_flatbuffers);
criterion_main!(benches);
