package selectors

import (
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/kva3umoda/go-ga/genome"
)

// Турнирный отбор (Tournament selection)
// Турнирный отбор может быть описан следующим образом: из популяции, содержащей N строк, выбирается случайным
// образом t строк и лучшая строка записывается в промежуточный массив (между выбранными строками проводится турнир).
// Эта операция повторяется N раз. Строки в полученном промежуточном массиве затем используются для скрещивания
// (также случайным образом). Размер группы строк, отбираемых для турнира часто равен 2. В этом случае говорят
// о двоичном/парном турнире (binary tournament). Вообще же t называется численностью турнира (tournament size).
// Чем больше турнир, тем более жесткий вариант селекции, т.е. тем меньше шансов у особей, "кому за, или кто просто плохо".
//
// Преимуществом данной стратегии является то, что она не требует дополнительных вычислений и упорядочивания строк в популяции
// по возрастанию приспособленности. Также, на мой взгляд, такой вариант селекции ближе к реальности, т.к. успешность той
// или иной особи во многом определяется ее окружением, насколько оно лучше или хуже ее. А иначе получилось бы, что бац!
// появился где-нибудь в Бразилии супер-таракан, а в России все усатые собратья тараканьи взяли (тоже бац!) и вымерли,
// от осознания своей неприспособленности. ;)
type Tournament struct {
	rnd       *rand.Rand
	lock      sync.Mutex
	tournSize int // Размер турнира
}

func NewTournament(tournSize int) *Tournament {
	return &Tournament{
		rnd:       rand.New(rand.NewSource(time.Now().UnixNano())),
		tournSize: tournSize,
	}
}

func (r *Tournament) Select(populationSize int, individuals []genome.Individual) []genome.Individual {
	chosen := make([]genome.Individual, 0, populationSize)

	for i := 0; i < populationSize; i++ {
		// смешиваем
		r.rnd.Shuffle(len(individuals), func(i, j int) {
			individuals[i], individuals[j] = individuals[j], individuals[i]
		})
		// Сортировка по возрастанию
		sort.Slice(individuals[0:r.tournSize], func(i, j int) bool {
			return individuals[i].Fitness > individuals[j].Fitness
		})
		// берем самого лучшего
		chosen = append(chosen, individuals[0].Clone())
	}

	return chosen
}

/**
def selTournament(individuals, k, tournsize, fit_attr="fitness"):
    """Select the best individual among *tournsize* randomly chosen
    individuals, *k* times. The list returned contains
    references to the input *individuals*.

    :param individuals: A list of individuals to select from.
    :param k: The number of individuals to select.
    :param tournsize: The number of individuals participating in each tournament.
    :param fit_attr: The attribute of individuals to use as selection criterion
    :returns: A list of selected individuals.

    This function uses the :func:`~random.choice` function from the python base
    :mod:`random` module.
    """
    chosen = []
    for i in xrange(k):
        aspirants = selRandom(individuals, tournsize)
        chosen.append(max(aspirants, key=attrgetter(fit_attr)))
    return chosen
*/
