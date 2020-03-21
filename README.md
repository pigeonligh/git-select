# git-select
Select user for git

### Installation

```
go install github.com/pigeonligh/git-select/cmd/git-select
```

or 

```
go install github.com/pigeonligh/git-select/cmd/gsel
```

### Usage

Add config:

```
git-select add
```

List configs:

```
git-select list
```

Remove config:

```
git-select rm config-tag
```

Select config:

```
git-select -t config-tag
```

or (for global)

```
git-select -t config-tag -g
```
