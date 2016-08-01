package tests

import (
    "github.com/sadhanandh/revelbasic"
    "github.com/sadhanandh/revelbasic/application/app/models"
    "github.com/revel/revel/testing"
)

type AppTest struct {
    testing.TestSuite
}

func (t *AppTest) Before() {
    // Make sure our collection is clean
    models.Collection(revelbasic.Session).DropCollection()
}

func (t *AppTest) TestThatIndexPageWorks() {
    t.Get("/")
    t.AssertOk()
    t.AssertContentType("text/html; charset=utf-8")

}

func (t *AppTest) TestSave() {
    b := models.GetBook("MobyDick")
    t.AssertEqual("Moby Dick", b.Title)
    b.Save(revelbasic.Session)
    d := models.GetBookByObjectId(revelbasic.Session, b.Id)
    t.AssertEqual(b.Title, d.Title)
    t.AssertEqual(b.Id, d.Id)
    t.AssertEqual(b.Body, d.Body)
    t.AssertEqual(b.Tags, d.Tags)
}

func (t *AppTest) After() {
    // Cleanup any mess we made
    models.Collection(revelbasic.Session).DropCollection()
}
