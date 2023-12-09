string[] input = File.ReadAllLines("./input.txt");

var numMapping = new Dictionary<string, char>() {
    { "zero", '0' },
    { "one", '1' },
    { "two", '2' },
    { "three", '3' },
    { "four", '4' },
    { "five", '5' },
    { "six", '6' },
    { "seven", '7' },
    { "eight", '8' },
    { "nine", '9' },
};

Console.WriteLine($"Part 1: {PartOne()}");
Console.WriteLine($"Part 2: {PartTwo()}");

int PartOne() {
    var calibrationValues = new List<int>();

    foreach (string line in input) {
        if (string.IsNullOrEmpty(line)) {
            continue;
        }

        char? firstDigit = null;
        char? lastDigit = null;

        foreach (char c in line) {
            if (int.TryParse(c.ToString(), out int digit)) {
                if (firstDigit == null) {
                    firstDigit = c;
                }

                if (lastDigit == null) {
                    lastDigit = firstDigit;
                } else {
                    lastDigit = c;
                }
            }
        }

        string num = $"{firstDigit}{lastDigit}";
        calibrationValues.Add(int.Parse(num));
    }

    return calibrationValues.Sum();
}

int PartTwo() {
    var calibrationValues = new List<int>();

    foreach (string line in input) {
        if (string.IsNullOrEmpty(line)) {
            break;
        }

        var dict = new Dictionary<int, char>();

        foreach (KeyValuePair<string, char> entry in numMapping) {
            if (line.Contains(entry.Key)) {
                var indexes = line.AllIndexesOf(entry.Key);

                foreach (var index in indexes) {
                    dict.Add(index, char.Parse(entry.Value.ToString()));
                }

            } 
            
            if (line.Contains(entry.Value)) {
                var indexes = line.AllIndexesOf(entry.Value.ToString());

                foreach (var index in indexes) {
                    dict.Add(index, char.Parse(entry.Value.ToString()));
                }
            }
        }

        var firstDigit = dict[dict.Select(x => x.Key).ToArray().Min()];
        var lastDigit = dict[dict.Select(x => x.Key).ToArray().Max()];

        string num = $"{firstDigit}{lastDigit}";
        calibrationValues.Add(int.Parse(num));
    }

    return calibrationValues.Sum();
}

public static class Ext {
    public static IEnumerable<int> AllIndexesOf(this string str, string searchString) {
        int minIndex = str.IndexOf(searchString);
        while (minIndex != -1) {
            yield return minIndex;
            minIndex = str.IndexOf(searchString, minIndex + searchString.Length);
        }
    }
}
