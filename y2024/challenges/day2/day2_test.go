package day2

import (
    "testing"
)

func TestIsReportSafe(t *testing.T) {
    // report1 := []int{7, 6, 4, 2, 1}
    // report2 := []int{1, 2, 7, 8, 9}
    // report3 := []int{9, 7, 6, 2, 1}
    report4 := []int{1, 3, 2, 4, 5}
    // report5 := []int{8, 6, 4, 4, 1}
    // report6 := []int{1, 3, 6, 7, 9}
    // report7 := []int{1, 3, 6, 7, 10}
    // report8 := []int{1, 9, 2, 3}

    // Part 1
    // if isReportSafe(report1) != true {
    //    t.Errorf("Report expected to be safe")
    // }
    // if isReportSafe(report2) != false {
    //    t.Errorf("Report expected to be not safe")
    // }
    // if isReportSafe(report3) != false {
    //        t.Errorf("Report expected to be not safe")
    // }
    // if isReportSafe(report4) != false {
    //        t.Errorf("Report expected to be not safe")
    // }
    // if isReportSafe(report5) != false {
    //        t.Errorf("Report expected to be not safe")
    // }
    // if isReportSafe(report6) != true {
    //        t.Errorf("Report expected to be safe")
    // }

    // Part 2
    // if isReportSafe(report1) != true {
    //    t.Errorf("Report expected to be safe")
    // }
    // if isReportSafe(report2) != false {
    //        t.Errorf("Report expected to not be safe")
    // }
    // if isReportSafe(report3) != false {
    //        t.Errorf("Report expected to not be safe")
    // }
    if isReportSafe(report4) != true {
           t.Errorf("Report expected to be not safe")
    }
    // if isReportSafe(report5) != true {
    //        t.Errorf("Report expected to be safe")
    // }
    // if isReportSafe(report6) != true {
    //        t.Errorf("Report expected to be safe")
    // }
    // if isReportSafe(report7) != true {
    //        t.Errorf("Report expected to be safe")
    // }
    // if isReportSafe(report8) != true {
    //        t.Errorf("Report expected to be safe")
    // }
}
