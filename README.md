# Boardgame-Bibliothek

In diesem Projekt wird eine Basis-Datenstruktur für schachbrettartige Spielfelder
entwickelt. Darauf aufbauend entsteht eine Reihe von Funktionen, die für
unterschiedliche Spiele Bedingungen überprüfen und Spielzüge durchführen können.
Außerdem soll das Spielfeld auf der Konsole ausgegeben werden können.

Mit Hilfe diese Funktionen können anschließend kleinere Spiele oder Logikrätsel
implementiert werden.

## Aufgaben-Übersicht

Die Packages in diesem Repo bauen auf einander auf und sollten in der
folgenden Reihenfolge implementiert werden:

1. `board`: Implementierung einer Basis-Datenstruktur für schachbrettartige
   Spielfelder.
2. `queens`: Lösung des n-Damen-Problems auf einem Schachbrett.
3. `knights`: Lösung des Springer-Problems auf einem Schachbrett.

**Hinweise:**
Die Aufgaben in den Unter-Packages `queens/solvequeens` und `knights/solveknights` sollten mittels Rekursion gelöst werden und sind
komplexe Beispiele für sog. "Backtracking"-Algorithmen.

Die Aufgaben in den Basis-Packages `board`, `knights` und `queens` sind hingegen ohne Rekursion und mit Standard-Wissen zu Structs, Slices und Schleifen lösbar.
