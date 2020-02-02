"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest

import s1ap_types

from integ_tests.s1aptests import s1ap_wrapper


class TestSctpAbortAfterIdentityReq(unittest.TestCase):

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_sctp_abort_after_identity_req(self):
        """ testing SCTP Abort after Identity Request for a single UE """
        self._s1ap_wrapper.configUEDevice(1)

        req = self._s1ap_wrapper.ue_req
        print("************************* Running SCTP Abort after Identity"
              " Request for single UE for UE id ", req.ue_id)

        attach_req = s1ap_types.ueAttachRequest_t()
        attach_req.ue_Id = req.ue_id
        sec_ctxt = s1ap_types.TFW_CREATE_NEW_SECURITY_CONTEXT
        id_type = s1ap_types.TFW_MID_TYPE_GUTI
        eps_type = s1ap_types.TFW_EPS_ATTACH_TYPE_EPS_ATTACH
        attach_req.mIdType = id_type
        attach_req.epsAttachType = eps_type
        attach_req.useOldSecCtxt = sec_ctxt
        print("Sending Attach Request ue-id", req.ue_id)
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_ATTACH_REQUEST, attach_req)

        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertTrue(response, s1ap_types.tfwCmd.UE_IDENTITY_REQ_IND.value)
        print("Received Identity req ind ",
              s1ap_types.tfwCmd.UE_IDENTITY_REQ_IND.value)

        print("send SCTP ABORT")
        sctp_abort = s1ap_types.FwSctpAbortReq_t()
        sctp_abort.cause = 3
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.SCTP_ABORT_REQ, sctp_abort)


if __name__ == "__main__":
    unittest.main()
