CREATE TABLE mahasiswa (
    id_mahasiswa INT NOT NULL AUTO_INCREMENT,
    nama VARCHAR(100) NOT NULL,
    usia INT NOT NULL,
    gender INT NOT NULL,
    tanggal_registrasi DATETIME NOT NULL,
    PRIMARY KEY (id_mahasiswa)
);

CREATE TABLE jurusan (
    id_jurusan INT NOT NULL AUTO_INCREMENT,
    nama VARCHAR(100) NOT NULL,
    PRIMARY KEY (id_jurusan)
);

CREATE TABLE hobi (
    id_hobi INT NOT NULL AUTO_INCREMENT,
    nama VARCHAR(100) NOT NULL,
    PRIMARY KEY (id_hobi)
);

CREATE TABLE mahasiswa_hobi (
    id_mahasiswa INT NOT NULL,
    id_hobi INT NOT NULL,
    FOREIGN KEY (id_mahasiswa) REFERENCES mahasiswa(id_mahasiswa),
    FOREIGN KEY (id_hobi) REFERENCES hobi(id_hobi)
);

CREATE TABLE mahasiswa_jurusan (
    id_mahasiswa INT NOT NULL,
    id_jurusan INT NOT NULL,
    FOREIGN KEY (id_mahasiswa) REFERENCES mahasiswa(id_mahasiswa),
    FOREIGN KEY (id_jurusan) REFERENCES jurusan(id_jurusan)
)