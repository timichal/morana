/*
placing non-tile objects (the player, monsters, traps...) on the map
*/
package main

func placePlayer(floorid string, tile rune) {
    floor := gameMap.floorSet[floorid]
    // positioning the player to stairs down
    for i, row := range floor {
        for j := range row {
            if floor[i][j].TileType == tile {
                player.coord.X = i
                player.coord.Y = j
                return
            }
        }
    }
}