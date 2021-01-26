# Fed

Fed prints your money. Or at least your invoices, which is just future money.

Creates your invoices in a yaml schema and generates (or regenerates) PDFs as you
need.

Super green but working code is better than nothing.

## Installing

Compile the binary:

```
./scripts/compile.sh
```

Then move it somewhere on your PATH.

## Usage

Create a directory (or git repo) and add a `config.yml` following the example in
`./example/config.yml` (Theres not much validation yet).

Use `fed new` to generate invoices. The first argument is the client key in
the `config.yml`, further arguments should be a series of item keys and the
quantity of associated items to bill. e.g.:

```
fed new clientify cool-thing:4 other-thing:2
```

This will:
- Create a directory based on the current data to store the invoice (e.g.
  2021/Q1).
- Create a yaml file in the created directory representing the invoice and
  containing all the information necessary to generate a PDF.
- Create a PDF invoice in the created directory.
- Create a file tracking the last ID an invoice was created with. The number will
  be increased on creating subsequent invoices.

## e-Invoicing

Doesn't do it.

## TODO

- [ ] A lot of tidying up and documentation and CI and tests (I've been using the
      pdf like a repl...).
- [ ] Make the default font more platform agnostic

