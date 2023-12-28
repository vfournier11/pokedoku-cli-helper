# pokedoku CLI helper

Simple CLI helper for [pokedoku](https://pokedoku.com)

## Usage

Use the CLI helper to filter the list of Pokemon by ability, or type.

```bash
  go run main.go filter t:grass
  go run main.go filter a:overgrow t:grass
  go run main.go solve t:grass,t:fire,g:1 t:water,t:ground,t:flying
```

## References

- [pokedoku](https://pokedoku.com)
- [pokeapi](https://pokeapi.co/docs/v2)
