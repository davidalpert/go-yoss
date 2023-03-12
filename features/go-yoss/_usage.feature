Feature: usage
  Background:
    Given I have installed "yoss" locally into the path

  # @announce-stdout @announce-stderr
  Scenario: help
    When I run `yoss`
    Then the stdout should show usage