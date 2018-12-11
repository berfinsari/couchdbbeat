package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/berfinsari/couchdbbeat/config"
)

// Couchdbbeat configuration.
type Couchdbbeat struct {
	done     chan struct{}
	config   config.CouchdbbeatConfig
	client   beat.Client
	CbConfig config.ConfigSettings
	period   time.Duration
	port     string
	host     string
}

const selector = "couchdbbeat"

// New creates an instance of couchdbbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	cb := &Couchdbbeat{
		done: make(chan struct{}),
	}
	err := cfgfile.Read(&cb.CbConfig, "")
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return nil, fmt.Errorf("Error reading configuration file: %v", err)
	}
	return cb, nil
}

// Run starts couchdbbeat.
func (cb *Couchdbbeat) Run(b *beat.Beat) error {
	logp.Info("couchdbbeat is running! Hit CTRL-C to stop it.")
	cb.CheckConfig(b)

	var err error
	cb.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(cb.period)
	for {
		select {
		case <-cb.done:
			return nil
		case <-ticker.C:
		}
		serverstats, err := cb.getServerStats(b)
		if err != nil {
			logp.Debug(selector, "Error reading server stats")
			return err
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":   b.Info.Name,
				"server": serverstats,
			},
		}
		cb.client.Publish(event)
		logp.Info("Event sent")
	}
}

func (cb *Couchdbbeat) CheckConfig(b *beat.Beat) error {
	if cb.CbConfig.Input.Period != nil {
		cb.period = time.Duration(*cb.CbConfig.Input.Period) * time.Second
	} else {
		cb.period = 30 * time.Second
	}

	if cb.CbConfig.Input.Port != nil {
		cb.port = *cb.CbConfig.Input.Port
	} else {
		cb.port = "5984"
	}

	if cb.CbConfig.Input.Host != nil {
		cb.host = *cb.CbConfig.Input.Host
	} else {
		cb.port = "localhost"
	}

	logp.Debug(selector, "Init Couchdbbeat")
	logp.Debug(selector, "Port %v", cb.port)
	logp.Debug(selector, "Period %v", cb.period)
	logp.Debug(selector, "Host %v", cb.host)

	return nil
}

// Stop stops couchdbbeat.
func (cb *Couchdbbeat) Stop() {
	cb.client.Close()
	close(cb.done)
}
