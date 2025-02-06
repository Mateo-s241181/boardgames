# N-Damen-Problem

In diesem Package wird das n-Damen-Problem auf einem Schachbrett gelöst.

Die Basis bildet der `Board`-Datentyp aus dem Package `board`.
Dieser Datentyp stellt bereits ein schachbrettartiges Spielfeld dar und bietet Funktionen, um zu prüfen, ob eine Figur auf einer Zeile, Spalte oder Diagonale vorhanden bzw. nicht vorhanden ist.

Darauf aufbauend wird hier in `queenlogic.go` eine Funktion implementiert, die prüft, ob nach den Schach-Regeln eine Dame auf einem Schachbrett platziert werden kann, ohne eine andere Dame anzugreifen.

Im Unter-Package `solvequeens` wird schließlich eine Funktion implementiert, die das n-Damen-Problem auf einem Schachbrett löst.
