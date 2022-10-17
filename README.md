# Leftover binary cleaner

When compiled in debug mode, Rust binaries are quite bloated. As one bloats up their projects directory with random ideas that didn't go anywhere, a significant amount
of harddrive space is lost to binaries that are never run. There's no reason one should have to manually type "cargo clean" on a bunch of projects to reclaim some disk space.

This project is a quick little utility that currently enumerates all Rust projects (defined as a folder with a `Cargo.toml`) and deletes the `target` directory. It will only try to
delete `target` if it sees a `Cargo.toml`.

I just whipped this up in an afternoon, however I may flesh this out to be more language agnostic. In addition, I would also like to refine this little tool so it prints how much
space it freed up.
