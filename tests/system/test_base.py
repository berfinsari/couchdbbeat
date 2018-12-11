from couchdbbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Couchdbbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        couchdbbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("couchdbbeat is running"))
        exit_code = couchdbbeat_proc.kill_and_wait()
        assert exit_code == 0
