# Package für Spielfelder

## Empfohlene Aufgaben-Reihenfolge

1. `listtools.go`:
    Implementierung von Hilfsfunktionen für Listen,
    die z.B. prüfen können, ob ein Element in einer Liste enthalten ist.
2. `boardtools.go`:
    Implementierung von Hilfsfunktionen für Spielfelder,
    die z.B. prüfen können, ob eine Zeile oder Spalte eines Spielfelds
    nur aus einem bestimmten Element besteht oder ein bestimmtes Element
    Diese Funktionen sind allgemein nützlich, um den Zustand eines Spielfelds abzufragen.
    (nicht) enthält.
3. `boardlogic_chess.go`:
    Implementierung von Funktionen für ein Schachbrett,
    die z.B. prüfen können, ob eine Dame auf einer Zeile, Spalte oder
    Diagonale vorhanden ist.
    Diese Funktionen sind spezifisch für ein Schachspiel bzw. Schachprobleme.

Hinweis: Die Datei `output.go` enthält Funktionen, die das Spielfeld auf der Konsole ausgeben können.
Sie enthält aber keine Aufgaben, diese Funktionen sind bereits implementiert.
