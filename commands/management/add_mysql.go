// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package management

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/percona/pmm/api/managementpb/json/client"
	mysql "github.com/percona/pmm/api/managementpb/json/client/my_sql"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
)

var addMySQLResultT = commands.ParseTemplate(`
MySQL Service added.
Service ID  : {{ .Service.ServiceID }}
Service name: {{ .Service.ServiceName }}
`)

type addMySQLResult struct {
	Service *mysql.AddOKBodyService `json:"service"`
}

func (res *addMySQLResult) Result() {}

func (res *addMySQLResult) String() string {
	return commands.RenderTemplate(addMySQLResultT, res)
}

type addMySQLCommand struct {
	AddressPort string
	ServiceName string
	Username    string
	Password    string
}

func (cmd *addMySQLCommand) Run() (commands.Result, error) {
	status, err := agentlocal.GetStatus()
	if err != nil {
		return nil, err
	}

	host, portS, err := net.SplitHostPort(cmd.AddressPort)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portS)
	if err != nil {
		return nil, err
	}

	params := &mysql.AddParams{
		Body: mysql.AddBody{
			PMMAgentID:  status.AgentID,
			NodeID:      status.NodeID,
			ServiceName: cmd.ServiceName,
			Address:     host,
			Port:        int64(port),

			MysqldExporter: true,
			Username:       cmd.Username,
			Password:       cmd.Password,

			QANUsername:        cmd.Username,
			QANPassword:        cmd.Password,
			QANMysqlPerfschema: true,
		},
		Context: commands.Ctx,
	}
	resp, err := client.Default.MySQL.Add(params)
	if err != nil {
		return nil, err
	}

	return &addMySQLResult{
		Service: resp.Payload.Service,
	}, nil
}

// register command
var (
	AddMySQL  = new(addMySQLCommand)
	AddMySQLC = AddC.Command("mysql", "Add MySQL to monitoring.")
)

func init() {
	AddMySQLC.Arg("address", "MySQL address and port. Default: 127.0.0.1:3306.").Default("127.0.0.1:3306").StringVar(&AddMySQL.AddressPort)

	hostname, _ := os.Hostname()
	serviceName := hostname + "-mysql"
	serviceNameHelp := fmt.Sprintf("Service name. Default: %s.", serviceName)
	AddMySQLC.Arg("name", serviceNameHelp).Default(serviceName).StringVar(&AddMySQL.ServiceName)

	AddMySQLC.Flag("username", "MySQL username.").StringVar(&AddMySQL.Username)
	AddMySQLC.Flag("password", "MySQL password.").StringVar(&AddMySQL.Password)
}
