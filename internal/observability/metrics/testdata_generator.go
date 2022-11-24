package metrics

import (
	"math/rand"
	"time"

	"github.com/jmcvetta/randutil"
)

const maxTestCommandDuration = 1000 // ms

const maxTestCommandTimeInterval = 10 // ms

const maxUpdateTestCommandWeightsTimeInterval = 1000 // ms

var testCommands = []string{"/test_1", "/test_2", "/test_3", "/test_4", "/test_5"}

// testDataGenerator искуственно генерирует большое количество данных для метрик
//  в целях тестирования инфраструктуры сбора метрик (Prometheus, Grafana).
type testDataGenerator struct {
	rnd      *rand.Rand
	commands []randutil.Choice
}

// newTestDataGenerator создаёт генератор тестовых данных для метрик.
func newTestDataGenerator() (*testDataGenerator, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixMicro()))

	commands := make([]randutil.Choice, 0, len(testCommands))
	for _, command := range testCommands {
		commands = append(commands, randutil.Choice{Item: command, Weight: rnd.Intn(maxTestCommandDuration)})
	}

	return &testDataGenerator{
		rnd:      rnd,
		commands: commands,
	}, nil
}

// StartTestDataGenerator запускает генератор тестовых данных для метрик.
func StartTestDataGenerator() error {
	generator, err := newTestDataGenerator()
	if err != nil {
		return err
	}
	generator.start()
	return nil
}

// start запускает генератор тестовых данных для метрик.
func (g *testDataGenerator) start() {
	// Генерация метрик
	go func() {
		for {
			command, maxDuration := g.getRandomCommand()
			IncomingMessagesCount.Inc(command)
			IncomingMessageResponseTime.Set(command, time.Duration(rand.Intn(maxDuration))*time.Millisecond)
			time.Sleep(time.Duration(g.rnd.Intn(maxTestCommandTimeInterval)) * time.Millisecond)
		}
	}()

	// Обновление трендов графиков
	go func() {
		for {
			for i := 0; i < len(g.commands); i++ {
				if g.rnd.Intn(len(g.commands)) == 0 {
					command := &g.commands[i]
					command.Weight = g.rnd.Intn(maxTestCommandDuration)
				}
			}
			updateInterval := (maxUpdateTestCommandWeightsTimeInterval + g.rnd.Intn(maxUpdateTestCommandWeightsTimeInterval)) / 2
			time.Sleep(time.Duration(updateInterval) * time.Millisecond)
		}
	}()
}

// getRandomCommand возвращает случайную команду, в соответствии с весами команд.
func (g *testDataGenerator) getRandomCommand() (commandName string, maxDuration int) {
	command, err := randutil.WeightedChoice(g.commands)
	if err != nil {
		command = g.commands[g.rnd.Intn(len(g.commands))]
	}
	return command.Item.(string), command.Weight
}
