# aStar algorithm

Simple pathfinding algorithm in Go

### Configurations
You can configure initial state by `config.json`

```json
{
  "width": 10,
  "height": 10,
  "start": {
    "x": 2,
    "y": 4,
    "state": 3
  },
  "end": {
    "x": 8,
    "y": 5,
    "state": 3
  },
  "obstacles": [
    {
      "x": 5,
      "y": 3,
      "state": 1
    },
    {
      "x": 5,
      "y": 4,
      "state": 1
    },
    {
      "x": 5,
      "y": 5,
      "state": 1
    },
    {
      "x": 5,
      "y": 6,
      "state": 1
    }
  ]
}
```
Set `width`, `height` and `obstacles`

Each point can be in 5 states:
-	FREE (0) - free space
-	BLOCKED (1) - obstacles
-	START (2) - start point
-	END (3) - end point
-	PATH (4) - resust of path searching

### Results
After execution you get `out.png`

![alt text](https://raw.githubusercontent.com/JavaDar/aStar/master/out.png)
