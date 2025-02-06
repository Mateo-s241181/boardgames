# Boardgame-Bibliothek

In diesem Projekt wird eine Basis-Datenstruktur für schachbrettartige Spielfelder
entwickelt. Darauf aufbauend entsteht eine Reihe von Funktionen, die für
unterschiedliche Spiele Bedingungen überprüfen und Spielzüge durchführen können.
Außerdem soll das Spielfeld auf der Konsole ausgegeben werden können.

Mit Hilfe diese Funktionen können anschließend kleinere Spiele oder Logikrätsel
implementiert werden.

## Aufgaben-Übersicht

Das Package `board` enthält eine Basis-Datenstruktur für schachbrettartige Spielfelder, die in den anderen Packages verwendet wird.

Sie sollten also zuerst die Aufgaben in `board` bearbeiten, bevor Sie mit den anderen Packages beginnen.
Die Reihenfolge der anderen Packages spielt keine Rolle,
da diese unabängig voneinander sind.

Innerhalb der Packages gibt es weitere Hinweise in den jeweiligen README-Dateien.

Die Aufgaben in den Unter-Packages `queens/solvequeens` und `knights/solveknights` sollten mittels Rekursion gelöst werden und sind
komplexe Beispiele für sog. "Backtracking"-Algorithmen.

Die Aufgaben in den Basis-Packages `board`, `knights` und `queens` sind hingegen ohne Rekursion und mit Standard-Wissen zu Structs, Slices und Schleifen lösbar.
