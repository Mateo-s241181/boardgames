# Springer-Problem

In diesem Package wird das n-Damen-Problem auf einem Schachbrett gelöst.

Die Basis bildet der `Board`-Datentyp aus dem Package `board`.
Dieser Datentyp stellt bereits ein schachbrettartiges Spielfeld dar.

Darauf aufbauend werden hier in `knight.go` eine Funktion implementiert, die für ein Feld die möglichen Springerzüge berechnet bzw. prüft, ob für gegebene Koordinaten ein Springer gesetzt werden kann.

Im Unter-Package `solveknights` wird schließlich eine Funktion implementiert, die das Springer-Problem auf einem Schachbrett löst.
