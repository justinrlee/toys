#!/usr/local/bin/python

import os
import json

artifact = {}

if "ARTIFACT_TYPE" in os.environ:
  artifact["type"] = os.environ["ARTIFACT_TYPE"]

if "ARTIFACT_REFERENCE" in os.environ:
  artifact["reference"] = os.environ["ARTIFACT_REFERENCE"]

if "ARTIFACT_ARTIFACTACCOUNT" in os.environ:
  artifact["artifactAccount"] = os.environ["ARTIFACT_ARTIFACTACCOUNT"]

if "ARTIFACT_NAME" in os.environ:
  artifact["name"] = os.environ["ARTIFACT_NAME"]

if "ARTIFACT_VERSION" in os.environ:
  artifact["version"] = os.environ["ARTIFACT_VERSION"]

if "ARTIFACT_PROVENANCE" in os.environ:
  artifact["provenance"] = os.environ["ARTIFACT_PROVENANCE"]

if "ARTIFACT_METADATA" in os.environ:
  artifact["metadata"] = os.environ["ARTIFACT_METADATA"]

if "ARTIFACT_LOCATION" in os.environ:
  artifact["location"] = os.environ["ARTIFACT_LOCATION"]

if "ARTIFACT_UUID" in os.environ:
  artifact["uuid"] = os.environ["ARTIFACT_UUID"]

spinnaker_config = {
  "artifacts": [
    artifact
  ]
}

print("SPINNAKER_CONFIG_JSON=" + json.dumps(spinnaker_config))